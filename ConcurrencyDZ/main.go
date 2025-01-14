package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	randNumChan := make(chan []int)
	resultChan := make(chan []int)
	go randomNum(randNumChan)
	randNum := <-randNumChan
	go exponentiation(resultChan, randNum)
	fmt.Println(<-resultChan)
}

func randomNum(ch chan []int) {
	var randomSliceNums []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		randomSliceNums = append(randomSliceNums, rand.Intn(101))
	}
	ch <- randomSliceNums
}

func exponentiation(ch chan []int, num []int) {
	results := make([]int, len(num))
	for i, n := range num {
		results[i] = int(math.Pow(float64(n), 2))
	}
	ch <- results
}
