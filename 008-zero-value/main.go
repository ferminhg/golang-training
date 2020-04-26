package main

import "fmt"

var y string

func main() {
	printY()
	y = "Muken"
	printY()
}

func printY() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}