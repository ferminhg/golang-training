package main

import "fmt"

func main() {
	fmt.Println("gowopwop")
	// Exercise 1
	var x [5]int
	x[0] = 1
	x[1] = 10
	x[2] = 2
	x[3] = 20
	x[4] = 3

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)

	// Exercise 2
	y := [10]int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range y {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", y)
	// exercise 3
	a := y[:5]
	b := y[5:]
	c := y[2:7]
	d := y[1:6]
	fmt.Println(a, b, c, d)

	// Exercise 4
	x4 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x4 = append(x4, 52)
	fmt.Println(x4)
	x4 = append(x4, 53, 54, 55)
	fmt.Println(x4)
	y4 := []int{56, 57, 58, 59, 60}
	x4 = append(x4, y4...)
	fmt.Println(x4)

	// exercise 5
	x5 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y5 := append(x5[:3], x5[6:]...)
	fmt.Println(y5)

	// Exercise 6
	states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	statesUsa := make([]string, 50, 500)
	statesUsa = append(statesUsa, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Println(states)
	fmt.Println(len(statesUsa), cap(statesUsa))

	// Exercise 7
	stupidNames1 := []string{"James", "Bond", "Shaken, not stirred"}
	stupidNames2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	stupidNames := [][]string{stupidNames1, stupidNames2}

	for _, xs := range stupidNames {
		for _, val := range xs {
			fmt.Println(val)
		}
	}

	// Exercise 8
	mapNames := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(mapNames)
	for name, things := range mapNames {
		for _, thing := range things {
			fmt.Println(name, thing)
		}
	}

	// Exercise 9
	mapNames[`fermin`] = []string{`wop`, `wop`, `wop`}
	fmt.Println(mapNames)

	delete(mapNames, `fermin`)
	fmt.Println(mapNames)
}
