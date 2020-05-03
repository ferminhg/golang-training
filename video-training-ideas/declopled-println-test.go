package main

import "fmt"

func main() {
	wop(fmt.Println)
}

func wop(f func(a ...interface{}) (int, error)) {
	f("wopwop")
}