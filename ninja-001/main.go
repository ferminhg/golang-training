package main

import "fmt"

var x int
var y string
var z bool

func main() {
	ninja1()
	ninja2()
	ninja3()
	ninja4()
	ninja5()
}

func ninja1() {
	fmt.Println("ninja1")
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(x, y, z)
}

func ninja2() {
	fmt.Println("ninja2")
	fmt.Println(x, y, z)
}

func ninja3() {
	fmt.Println("ninja3")
	x = 42
	y = "James Bond"
	z = true
	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	fmt.Println(s)
}

type muken int

func ninja4() {
	fmt.Println("ninja4")
	var t muken
	fmt.Println(t)
	fmt.Printf("%T\n", t)
	t = 42
	fmt.Println(t)
}

var y5 int

func ninja5() {
	fmt.Println("ninja4")
	var t muken
	fmt.Println(t)
	fmt.Printf("%T\n", t)
	t = 42
	fmt.Println(t)
	fmt.Printf("%T\n", t)
	
	y5 = int(t)
	fmt.Println(y5)
	fmt.Printf("%T\n", y5)
}
