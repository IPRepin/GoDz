package verify

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/jordan-wright/email"
	"godz/3-validation-api/config"
	"net/http"
	"net/smtp"
	"strings"
	"sync"
)

type VerifyHandler struct {
	*config.Config
	hashes map[string]string
	mu     sync.Mutex
}

type VerifyHandlerDeps struct {
	*config.Config
}

func NewVerifyHandler(mux *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		Config: deps.Config,
		hashes: make(map[string]string),
	}
	mux.HandleFunc("POST /send", handler.SendEmail)
	mux.HandleFunc("GET /verify/", handler.VerifyGet)
}

func generateHash() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}

func (handler *VerifyHandler) SendEmail(w http.ResponseWriter, r *http.Request) {
	var request EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	hash, err := generateHash()
	if err != nil {
		http.Error(w, "Failed to generate verification hash", http.StatusInternalServerError)
		return
	}

	handler.mu.Lock()
	handler.hashes[hash] = request.Email
	handler.mu.Unlock()

	e := email.NewEmail()
	e.From = fmt.Sprintf("Verification Service <%s>", handler.Auth.EmailUser)
	e.To = []string{request.Email}
	e.Subject = "Email Verification"
	verificationLink := fmt.Sprintf("http://localhost:8080/verify/%s", hash)
	e.Text = []byte(fmt.Sprintf("Click the link to verify your email: %s", verificationLink))

	err = e.Send(
		handler.Auth.EmailHost+":"+handler.Auth.EmailPort,
		smtp.PlainAuth("", handler.Auth.EmailUser, handler.Auth.EmailPass, handler.Auth.EmailHost),
	)

	if err != nil {
		http.Error(w, "Failed to send verification email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"status": "verification email sent"})
	if err != nil {
		http.Error(w, "Failed to send verification email", http.StatusInternalServerError)
	}
}

func (handler *VerifyHandler) VerifyGet(w http.ResponseWriter, r *http.Request) {
	hash := strings.TrimPrefix(r.URL.Path, "/verify/")
	if hash == "" {
		http.Error(w, "Missing verification hash", http.StatusBadRequest)
		return
	}

	handler.mu.Lock()
	defer handler.mu.Unlock()

	if _, exists := handler.hashes[hash]; exists {
		delete(handler.hashes, hash)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("true"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("false"))
	}
}
