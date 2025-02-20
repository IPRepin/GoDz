package auth

import (
	"fmt"
	"godz/5-order-api-auth/configs"
	"godz/5-order-api-auth/internal/user"
	"godz/5-order-api-auth/pkg/jwt"
	"godz/5-order-api-auth/pkg/middleware"
	"godz/5-order-api-auth/pkg/req"
	"godz/5-order-api-auth/pkg/res"
	"net/http"
)

type AuthHandler struct {
	configs.Config
	UserRepo *user.UserRepository
}

type AuthHandlerDeps struct {
	configs.Config
	UserRepo *user.UserRepository
}

func NewAuthHandler(mux *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	mux.HandleFunc("POST /auth", middleware.IsAuth(handler.Auth(), deps.Config))
}

func (handler *AuthHandler) Auth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		phone, ok := r.Context().Value(middleware.ContextPhoneKey).(string)
		if ok {
			fmt.Println(phone)
		}
		reqBody, err := req.HandleBody[AuthRequest](w, r)
		if err != nil {
			return
		}
		// Если передан код
		if reqBody.Code != "" {
			findUser, err := handler.UserRepo.FindByPhone(reqBody.Phone)
			if err != nil {
				http.Error(w, "Пользователь не найден", http.StatusNotFound)
				return
			}
			if findUser.SessionId != reqBody.SessionId {
				http.Error(w, "Неверный session_id", http.StatusUnauthorized)
				return
			}
			if reqBody.Code != "3569" {
				http.Error(w, "Неверный код", http.StatusUnauthorized)
				return
			}
			token, err := jwt.NewJWT(handler.Config.Auth.Secret).CreateToken(jwt.JWTData{
				Phone: findUser.Phone,
			})
			if err != nil {
				http.Error(w, "Ошибка генерации токена", http.StatusInternalServerError)
				return
			}
			res.JsonResponse(w, AuthResponse{Token: token}, http.StatusOK)
			return
		}
		// Этап авторизации
		findUser, err := handler.UserRepo.FindByPhone(reqBody.Phone)
		if err != nil {
			// Пользователь не найден – создаём нового
			findUser = user.NewUser(reqBody.Phone)
			findUser, err = handler.UserRepo.Create(findUser)
			if err != nil {
				http.Error(w, "Ошибка создания пользователя", http.StatusInternalServerError)
				return
			}
		} else {
			// Пользователь существует – обновляем session_id
			findUser.GenerateSessionId()
			findUser, err = handler.UserRepo.UpdateSessionId(findUser.SessionId, reqBody.Phone)
			if err != nil {
				http.Error(w, "Ошибка обновления session_id", http.StatusInternalServerError)
				return
			}
		}
		// Отправляем новый session_id для дальнейшей верификации кода
		res.JsonResponse(w, AuthResponse{SessionId: findUser.SessionId}, http.StatusOK)
	}
}
