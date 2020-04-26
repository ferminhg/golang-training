package main

import "fmt"

func main() {
	// https://godoc.org/fmt#example-Println
	// _ variable to ignore
	n, _  := fmt.Println("wopwop", 42, true)
	fmt.Println(n)
}
