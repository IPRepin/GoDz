package verify

import (
	"encoding/json"
	"fmt"
	"github.com/jordan-wright/email"
	"godz/3-validation-api/config"
	"godz/3-validation-api/pkg/res"
	"math/rand"
	"net/http"
	"net/smtp"
)

type VerifyHandler struct {
	*config.Config
}

type VerifyHandlerDeps struct {
	*config.Config
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

func (handler *VerifyHandler) SendEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request EmailRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		hash, err := generateHash()
		if err != nil {
			http.Error(w, "Failed to generate verification hash", http.StatusInternalServerError)
			return
		}
		e := email.NewEmail()
		e.From = handler.Config.Auth.EmailAddress
		e.To = []string{request.Email}
		e.Subject = "Verify your email"
		verificationLink := fmt.Sprintf("http://localhost:8080/verify/%s", hash)
		e.HTML = []byte(fmt.Sprintf("<h1>Click the link to verify your email: %s</h1>", verificationLink))
		err = e.Send(
			handler.Auth.EmailHost+":"+handler.Auth.EmailPort,
			smtp.PlainAuth("", handler.Auth.EmailUser, handler.Auth.EmailPass, handler.Auth.EmailHost),
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.JsonResponse(w, request, http.StatusOK)
	}
}

func (handler *VerifyHandler) VerifyGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
