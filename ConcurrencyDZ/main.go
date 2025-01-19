package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	randNumChan := make(chan int, 10)
	resultChan := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			randomNum(randNumChan)
		}
		close(randNumChan)
	}()
	var randSlice []int
	for randNum := range randNumChan {
		randSlice = append(randSlice, randNum)
	}
	go exponentiation(resultChan, randSlice)
	fmt.Println(<-resultChan)
}

func randomNum(ch chan int) {
	rand.Seed(time.Now().UnixNano())
	numRandom := rand.Intn(101)
	ch <- numRandom
}

func exponentiation(ch chan string, num []int) {
	var results []string
	for _, n := range num {
		results = append(results, strconv.Itoa(int(math.Pow(float64(n), 2))))
	}
	ch <- strings.Join(results, " ")
}
