package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
)

const routines = 100
const primesPerRoutes = 100

func main() {
	ch := sender()
	reciever(ch)
	fmt.Println("bye bye ✌️")
}

func sender() <-chan *big.Int {
	ch := make(chan *big.Int)
	var wg sync.WaitGroup
	go launchRoutines(&wg, ch)
	return ch
}

func launchRoutines(wg *sync.WaitGroup, ch chan<- *big.Int) {
	for i := 0; i < routines; i++ {
		wg.Add(1)
		go generatePrimes(wg, ch)
	}
	wg.Wait()
	close(ch)
}

func generatePrimes(wg *sync.WaitGroup, ch chan<- *big.Int) {
	for i := 0; i < primesPerRoutes; i++ {
		p, _ := rand.Prime(rand.Reader, 128)
		ch <- p
	}
	wg.Done()
}

func reciever(ch <-chan *big.Int) {
	position := 0
	for prime := range ch {
		fmt.Println(position, " -> ", prime)
		position++
	}
}
