package verify

import (
	"encoding/json"
	"github.com/jordan-wright/email"
	"godz/3-validation-api/config"
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
	handler := &VerifyHandler{deps.Config}
	mux.HandleFunc("POST /verify/send", handler.SendEmail)
	mux.HandleFunc("GET /verify/{hash}", handler.VerifyGet)
}

func (handler *VerifyHandler) SendEmail(w http.ResponseWriter, r *http.Request) {
	var request EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	e := email.NewEmail()
	e.From = "Ваше имя <" + handler.Config.Auth.EmailUser + ">"
	e.To = []string{request.To}
	e.Subject = request.Subject
	e.Text = []byte(request.Body)

	err := e.Send(
		handler.Config.Auth.EmailHost+":"+handler.Config.Auth.EmailPort,
		smtp.PlainAuth("", handler.Config.Auth.EmailUser,
			handler.Config.Auth.EmailPass, handler.Config.Auth.EmailHost),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *VerifyHandler) VerifyGet(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Query().Get("hash")
	if hash == "" {
		http.Error(w, "Hash is required", http.StatusBadRequest)
		return
	}
	w.Write([]byte("Email verified successfully"))
}
