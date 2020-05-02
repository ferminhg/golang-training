package main

import "fmt"

type person struct {
	first, last     string
	iceCreamFlavors []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	// exercise 1
	p1 := person{
		first:           "Fermin",
		last:            "Hdez",
		iceCreamFlavors: []string{"nata", "pistacho"},
	}
	p2 := person{
		first:           "Lucia",
		last:            "Mezquida",
		iceCreamFlavors: []string{"mantecado", "avellana"},
	}
	//fmt.Println(p1, p2)

	// exercise 2
	people := map[string]person{}
	people[p1.last] = p1
	people[p2.last] = p2

	fmt.Println(people)

	// exercise 3
	myTruck := truck{
		vehicle: vehicle{
			doors: 4,
			color: "green",
		},
		fourWheel: true,
	}

	mySedan := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: false,
	}
	fmt.Println("my truck:", myTruck.doors, myTruck.color, myTruck.fourWheel)
	fmt.Println("my sedan:", mySedan.doors, mySedan.color, mySedan.luxury)

	mybike := struct {
		color       string
		singleSpeed bool
	}{
		color:       "red",
		singleSpeed: true,
	}

	fmt.Println(mybike)

	netflixUser := struct {
		name       string
		favSeries  []string
		viewSeries map[string]int
	}{
		name: "Fermin",
		favSeries: []string{
			"Friends",
			"X-Files",
			"Game of Thrones",
		},
		viewSeries: map[string]int{
			"Friends":         9,
			"X-Files":         2,
			"Game of Thrones": 4,
		},
	}

	fmt.Println(netflixUser)
}
