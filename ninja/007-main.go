package main

import "fmt"

type person struct {
	first string
	last  string
}

func (p *person) changeMe(pNew person) {
	(*p).first = pNew.first
	(*p).last = pNew.last
}

func updateMe(p *person) {
	p.last = "Muka"
}


func main() {
	//exercise 1
	x := 82
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Printf("%F\n", &x)

	// exercise 2
	p1 := person {
		first: "Lucia",
		last: "Mezquida",
	}
	fmt.Println(p1)
	p1.changeMe(person{
		first: "Super",
		last: "programadora",
	} )
	fmt.Println(p1)
	updateMe(&p1)
	fmt.Println(p1)
}
