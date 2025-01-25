package verify

import (
	"encoding/json"
	"fmt"
	"github.com/jordan-wright/email"
	"godz/3-validation-api/config"
	"godz/3-validation-api/pkg/req"
	"godz/3-validation-api/pkg/res"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	"sync"
)

type VerifyHandler struct {
	*config.Config
	mu sync.Mutex
}

type VerifyHandlerDeps struct {
	*config.Config
}

type VerificationData struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

func NewVerifyHandler(mux *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		Config: deps.Config,
	}
	mux.HandleFunc("POST /send", handler.SendEmail())
	mux.HandleFunc("GET /verify/{hash}", handler.VerifyGet())
}

func generateHash() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}

func (handler *VerifyHandler) saveVerificationData(email, hash string) error {
	handler.mu.Lock()
	defer handler.mu.Unlock()

	// Чтение существующих данных из файла
	data, err := handler.readVerificationData()
	if err != nil {
		return err
	}

	// Добавление новой записи
	data = append(data, VerificationData{Email: email, Hash: hash})

	// Сериализация данных в JSON
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Запись данных в файл
	return os.WriteFile("hash.json", file, 0644)
}

// readVerificationData читает данные из файла hash.json
func (handler *VerifyHandler) readVerificationData() ([]VerificationData, error) {
	var data []VerificationData

	// Проверяем, существует ли файл
	if _, err := os.Stat("hash.json"); os.IsNotExist(err) {
		return data, nil // Файл не существует, возвращаем пустой список
	}

	// Чтение файла
	file, err := os.ReadFile("hash.json")
	if err != nil {
		return nil, err
	}

	// Десериализация данных
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// deleteVerificationData удаляет запись из файла hash.json
func (handler *VerifyHandler) deleteVerificationData(hash string) error {
	handler.mu.Lock()
	defer handler.mu.Unlock()

	// Чтение существующих данных из файла
	data, err := handler.readVerificationData()
	if err != nil {
		return err
	}

	// Удаление записи с указанным hash
	newData := make([]VerificationData, 0)
	for _, item := range data {
		if item.Hash != hash {
			newData = append(newData, item)
		}
	}

	// Сериализация обновленных данных в JSON
	file, err := json.MarshalIndent(newData, "", "  ")
	if err != nil {
		return err
	}

	// Запись обновленных данных в файл
	return os.WriteFile("hash.json", file, 0644)
}

func (handler *VerifyHandler) sendVerificationEmail(emailAddress, hash string) error {
	e := email.NewEmail()
	e.From = handler.Config.Auth.EmailAddress
	e.To = []string{emailAddress}
	e.Subject = "Verify your email"
	verificationLink := handler.Auth.UrlVerify + hash
	e.HTML = []byte(fmt.Sprintf("<h1>Click the link to verify your email: %s</h1>", verificationLink))

	err := e.Send(
		handler.Auth.EmailHost+":"+handler.Auth.EmailPort,
		smtp.PlainAuth("", handler.Auth.EmailUser, handler.Auth.EmailPass, handler.Auth.EmailHost),
	)
	if err != nil {
		return err
	}
	return nil
}

func (handler *VerifyHandler) SendEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandelBody[EmailRequest](&w, r)
		if err != nil {
			return
		}
		hash, err := generateHash()
		if err != nil {
			res.JsonResponse(w, "Failed to generate verification hash", http.StatusInternalServerError)
			return
		}
		if err := handler.sendVerificationEmail(body.Email, hash); err != nil {
			res.JsonResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Сохраняем данные в файл
		if err := handler.saveVerificationData(body.Email, hash); err != nil {
			res.JsonResponse(w, "Failed to save verification data", http.StatusInternalServerError)
			return
		}

		res.JsonResponse(w, body, http.StatusOK)
	}
}

func (handler *VerifyHandler) VerifyGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandelBody[VerifyRequest](&w, r)
		if err != nil {
			return
		}
	}
}
