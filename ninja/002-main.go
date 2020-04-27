package main

import "fmt"

func main() {
	level1()
	level2()
	level3()
	level4()
	level5()
	level6()
}

func level1() {
	x := 13
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
}

func level2() {
	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)
	fmt.Println(a, b, c, d, e, f)
}

const dolar = "$"
const monthReveneu float32 = 1.4

func level3() {
	fmt.Println(monthReveneu)
	fmt.Println(dolar)
}

func level4() {
	x := 4
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d\t%b\t%#x\n", y, y, y)
}

func level5() {
	x := `wopwop	wopwop`
	fmt.Println(x)
}

const (
	y1 = iota + 2020
	y2
	y3
	y4
)

func level6() {
	fmt.Println(y1, y2, y3, y4)
}
