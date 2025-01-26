package verify

import (
	"encoding/json"
	"fmt"
	"github.com/jordan-wright/email"
	"godz/3-validation-api/config"
	"godz/3-validation-api/pkg/req"
	"godz/3-validation-api/pkg/res"
	"log"
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

	data, err := handler.readVerificationData()
	if err != nil {
		return err
	}

	data = append(data, VerificationData{Email: email, Hash: hash})

	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(handler.File.FileName, file, 0644)
}

func (handler *VerifyHandler) readVerificationData() ([]VerificationData, error) {
	var data []VerificationData

	if _, err := os.Stat(handler.File.FileName); os.IsNotExist(err) {
		return data, nil
	}

	file, err := os.ReadFile(handler.File.FileName)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (handler *VerifyHandler) deleteVerificationData(hash string) error {
	handler.mu.Lock()
	defer handler.mu.Unlock()

	data, err := handler.readVerificationData()
	if err != nil {
		return err
	}

	newData := make([]VerificationData, 0)
	for _, item := range data {
		if item.Hash != hash {
			newData = append(newData, item)
		}
	}

	file, err := json.MarshalIndent(newData, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(handler.File.FileName, file, 0644)
}

func (handler *VerifyHandler) sendVerificationEmail(emailAddress, hash string) error {
	e := email.NewEmail()
	e.From = handler.Config.Auth.EmailAddress
	e.To = []string{emailAddress}
	e.Subject = "Verify your email"
	verificationLink := handler.Url.UrlVerify + hash
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
		body, err := req.HandleBody[EmailRequest](w, r)
		if err != nil {
			log.Printf("Ошибка обработки запроса: %v", err)
			return
		}

		hash, err := generateHash()
		if err != nil {
			log.Printf("Ошибка генерации хеша: %v", err)
			res.JsonResponse(w, map[string]string{"error": "Внутренняя ошибка сервера"}, http.StatusInternalServerError)
			return
		}

		if err := handler.sendVerificationEmail(body.Email, hash); err != nil {
			log.Printf("Ошибка отправки email: %v", err)
			res.JsonResponse(w, map[string]string{"error": "Ошибка отправки письма"}, http.StatusInternalServerError)
			return
		}

		if err := handler.saveVerificationData(body.Email, hash); err != nil {
			log.Printf("Ошибка сохранения данных: %v", err)
			res.JsonResponse(w, map[string]string{"error": "Внутренняя ошибка сервера"}, http.StatusInternalServerError)
			return
		}

		res.JsonResponse(w, map[string]string{"status": "Письмо отправлено"}, http.StatusOK)
	}
}

func (handler *VerifyHandler) VerifyGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")

		data, err := handler.readVerificationData()
		if err != nil {
			res.JsonResponse(w, "Failed to read verification data", http.StatusInternalServerError)
			return
		}

		found := false
		for _, item := range data {
			if item.Hash == hash {
				found = true
				break
			}
		}

		if found {
			if err := handler.deleteVerificationData(hash); err != nil {
				res.JsonResponse(w, "Failed to delete verification data", http.StatusInternalServerError)
				return
			}
		}

		res.JsonResponse(w, found, http.StatusOK)
	}
}
