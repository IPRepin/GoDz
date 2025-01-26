package res

import (
	"encoding/json"
	"log"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Ошибка формирования JSON ответа: %v", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
	}
}
