package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	s := make(chan string)
	x := make(chan int, 1)

	go func() { c <- 42 }()
	go func() { s <- "wopwop" }()

	// send info in main routine
	x <- 31415
	fmt.Println(<-x)
	// te recieve again info un routine i need to put info again
	x <- 31416
	go func() {
		counter := <-x
		fmt.Println(counter)
		counter++
		x <- counter
	}()

	fmt.Println(<-c, <-s, <-x)
}
