package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"godz/5-order-api-auth/internal/user"
	"time"
)

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) CreateToken(u *user.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"phone": u.Phone,
		"exp":   expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return secret, nil
}
