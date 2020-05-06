package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type usuario struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := usuario{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := usuario{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := usuario{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	usuarios := []usuario{u1, u2, u3}

	fmt.Println(usuarios)
	//  json.NewEncoder(os.Stdout).encode(v interface{})

	err := json.NewEncoder(os.Stdout).Encode(usuarios)
	if err != nil {
		fmt.Println("We did something wrong and here's the error:", err)
	}
}
