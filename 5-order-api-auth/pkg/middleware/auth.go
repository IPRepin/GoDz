package middleware

import (
	"context"
	"godz/5-order-api-auth/configs"
	"godz/5-order-api-auth/pkg/jwt"
	"net/http"
	"strings"
)

type key string

const (
	ContextPhoneKey key = "ContextPhoneKey"
)

func writeUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	_, err := w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
	if err != nil {
		return
	}
}

func IsAuth(next http.Handler, config configs.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			writeUnauthorized(w)
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).ParseToken(token)
		if !isValid {
			writeUnauthorized(w)
		}
		ctx := context.WithValue(r.Context(), ContextPhoneKey, data.Phone)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	}
}
