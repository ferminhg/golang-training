package main

import "fmt"

var y = 43
var z int

func main() {
	// short declaration operator
	x := 42 // DECLARE a variable and ASSIGN a VALUE
	fmt.Println(x, y)

	foo()
	fmt.Println(z)
}

func foo() {
	fmt.Println("again: ", y)
}
