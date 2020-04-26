package main

import "fmt"

var y = 420

// https://godoc.org/fmt
func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%x\t%b\n", y, y, y)
	s := fmt.Sprintf("%#x\t%x\t%b\n", y, y, y)
	fmt.Println(s)
}