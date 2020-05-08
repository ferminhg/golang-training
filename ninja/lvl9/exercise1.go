package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("\twopwop")
	wg.Add(1)
	go foo()
	wg.Add(1)
	go bar()
	wg.Wait()
	fmt.Println("\tended")

}

func foo() {
	fmt.Println("foo")
	wg.Done()
}

func bar() {
	fmt.Println("bar")
	wg.Done()
}
