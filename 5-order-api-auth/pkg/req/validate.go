package req

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"strings"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Validate(payload interface{}) error {
	if err := validate.Struct(payload); err != nil {
		var errMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errMessages = append(errMessages, getReadableError(err))
		}
		return fmt.Errorf(strings.Join(errMessages, "; "))
	}
	return nil
}

func getReadableError(err validator.FieldError) string {
	field := strings.ToLower(err.Field())

	switch err.Tag() {
	case "required":
		return fmt.Sprintf("Поле '%s' обязательно для заполнения", field)
	case "phone":
		return "Некорректный формат phone"
	case "min":
		return fmt.Sprintf("Минимальная длина для '%s' - %s символов", field, err.Param())
	case "max":
		return fmt.Sprintf("Максимальная длина для '%s' - %s символов", field, err.Param())
	default:
		log.Printf("Неизвестная ошибка валидации: %s", err.Tag())
		return "Некорректное значение поля"
	}
}
