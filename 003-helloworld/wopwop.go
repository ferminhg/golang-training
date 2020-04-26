package main

import "fmt"

func main() {
	fmt.Println("wopwop")
	foo()
	fmt.Println("blabla")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("Bye bye")
}

// control flow
// - sequence
// - loop; iteractive
// - conditional
