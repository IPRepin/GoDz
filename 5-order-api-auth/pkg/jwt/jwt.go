package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTData struct {
	Phone string
}

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) CreateToken(data JWTData) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"phone": data.Phone,
		"exp":   expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return secret, nil
}

func (j *JWT) ParseToken(tokenString string) (bool, *JWTData) {
	t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}
	phone, ok := t.Claims.(jwt.MapClaims)["phone"].(string)
	if !ok {
		return false, nil
	}
	return t.Valid, &JWTData{
		Phone: phone,
	}
}
