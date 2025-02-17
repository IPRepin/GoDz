package req

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

func Decode[T any](body io.ReadCloser) (T, error) {
	var request T

	if body == nil {
		log.Print("Пустое тело запроса")
		return request, fmt.Errorf("отсутствуют данные в запросе")
	}
	defer body.Close()

	if err := json.NewDecoder(body).Decode(&request); err != nil {
		log.Printf("Ошибка декодирования JSON: %v", err)
		return request, fmt.Errorf("некорректный формат данных")
	}

	return request, nil
}
