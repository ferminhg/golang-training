package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p *person) speak() {
	fmt.Println("Ey ho ! - ", p.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		first: "Fermin",
		age:   38,
	}

	saySomething(&p)
	p.speak()
}
