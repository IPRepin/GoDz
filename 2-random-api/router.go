package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type RandomNum struct{}

func RandomHandler(router *http.ServeMux) {
	handler := &RandomNum{}
	router.HandleFunc("/random", handler.RandomNumber())
}

func (h *RandomNum) RandomNumber() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rand.Seed(time.Now().UnixNano())
		numRandom := rand.Intn(7)
		fmt.Println(numRandom)
		w.Write([]byte(strconv.Itoa(numRandom)))
	}
}
