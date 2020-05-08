package main

import "fmt"

func main() {
	fmt.Println("wopwop")
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	recieve(c)
	fmt.Println("bye bye")
}

func recieve(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
