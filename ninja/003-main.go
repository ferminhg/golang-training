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
	for i := 1; i <= 100000; i++ {
		fmt.Println(i)
	}
}

func level2() {
	for i := 65; i <= 126; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%U %c\n", i, i)
		}
	}
}

func level3() {
	i := 1
	for i <= 37 {
		fmt.Println(i)
		i++
	}
}

func level4() {
	i := 1
	for {
		if i <= 37 {
			break
		}
		fmt.Println(i)
		i++
	}
}

func level5() {
	for i := 10; i <= 100; i++ {
		if (i % 4) == 0 {
			fmt.Println(i)
		}
	}
}

func level6() {
	a := 3.0
	switch a {
	case 3:
		fmt.Println("string")
	case 2.0:
		fmt.Println("Float")
	}

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
