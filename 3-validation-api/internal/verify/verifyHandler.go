package verify

import (
	"godz/3-validation-api/config"
	"godz/3-validation-api/pkg/res"
	"net/http"
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

func (handler *VerifyHandler) SendEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := EmailRequest{
			Email: "a@a.com",
		}
		res.JsonResponse(w, response, http.StatusOK)
	}
}

func (handler *VerifyHandler) VerifyGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
