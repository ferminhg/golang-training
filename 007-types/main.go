package main

import "fmt"

var y = 42
// DECLARED is TYPE string
var z = `wop 

"wopwop"`
func main() {
	fmt.Println(z)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
}
