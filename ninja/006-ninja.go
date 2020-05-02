package main

import (
	"fmt"
	"math"
)

func main() {
	defer bye()
	//exercise 7
	defer func() {
		fmt.Println("bye bye mother fucker anoymous")
	}()

	fmt.Println("wopwop")
	y := foo()
	fmt.Println(y)
	x, s := bar()
	fmt.Println(x, s)

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(foo2(nums...))
	exec4()
	exec5()

	f := func(name string) {
		fmt.Println("ey bro!!", name)
	}
	f("Vinnie")

	fmt.Println("callback", operate(sum, 40, 3))
	exec10()
}

func sum(a int, b int) int {
	return a + b
}

// exercise 9
func operate(f func(a int, b int) int, a int, b int) int {
	return f(a, b)
}

func foo() int {
	return 28
}

func bar() (int, string) {
	return 28, "wopwop"
}

// exercise 2
func foo2(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bye() {
	fmt.Println("bye bye mother fucker")
}

type guy struct {
	first string
	last  string
	age   int
}

func (g guy) speak() {
	fmt.Println("My name:", g.first, "my age:", g.age)
}

func exec4() {
	g := guy{
		first: "fermin",
		last:  "hdez",
		age:   38,
	}
	g.speak()
}

type square struct {
	long   float32
	weigth float32
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * 2
}

func (s square) area() float32 {
	return s.long * s.weigth
}

type shape interface {
	area() float32
}

func info(sh shape) {
	fmt.Println("shape area:", sh.area())
}

func exec5() {
	s := square{
		long:   1.1,
		weigth: 4.3,
	}

	c := circle{
		radius: 1.45,
	}

	fmt.Println("square area", s.area())
	fmt.Println("cricle area", c.area())

	info(c)
	info(s)
}

func exec10(){
	f := fooEnclouse()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func fooEnclouse() func() int{
	x := 1
	return func () int {
		x += x
		return x*x
	}
}