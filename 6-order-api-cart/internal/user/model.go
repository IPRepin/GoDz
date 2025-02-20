package user

import (
	"godz/6-order-api-cart/configs"
	"godz/6-order-api-cart/internal/order"
	"gorm.io/gorm"
	"math/rand"
)

type RunesConfig struct {
	configs.Config
}

type User struct {
	gorm.Model
	Phone     string `json:"phone" gorm:"unique"`
	SessionId string `json:"session_id" gorm:"uniqueIndex"`
	Orders    []order.Order
}

func NewUser(phone string) *User {
	user := &User{
		Phone:     phone,
		SessionId: RandSessionId(),
	}
	user.GenerateSessionId()
	return user
}

func (user *User) GenerateSessionId() {
	user.SessionId = RandSessionId()
}

func RandSessionId() string {
	letterRunes := configs.GetConfig().Runs.LetterRunes
	randRune := make([]rune, 11)
	for i := range randRune {
		randRune[i] = rune(letterRunes[rand.Intn(len(letterRunes))])
	}
	return string(randRune)
}
