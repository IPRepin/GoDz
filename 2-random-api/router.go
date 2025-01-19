package main

import (
	"math/rand"
	"net/http"
	"strconv"
)

type RandomNum struct {
	min int
	max int
}

func RandomHandler(router *http.ServeMux, handler *RandomNum) {
	router.HandleFunc("/random", handler.RandomNumber())
}

func (h *RandomNum) RandomNumber() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		numRandom := rand.Intn(h.max-h.min) + h.min
		if _, err := w.Write([]byte(strconv.Itoa(numRandom))); err != nil {
			http.Error(w, "Ошибка записи в ответ", http.StatusInternalServerError)
			return
		}
	}
}
