package main

import (
	"fmt"
	"sort"
)

type Guy struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func (p Guy) String() string {
	sayings := ""
	sort.Strings(p.Sayings)
	for _, said := range(p.Sayings) {
		sayings += fmt.Sprintf("\t- %s\n", said)
	}
	return fmt.Sprintf("%s %s: %d \n%s\n", p.First, p.Last, p.Age, sayings)
}

type ByAge []Guy

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []Guy

func (a ByLast) Len() int           { return len(a) }
func (a ByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

func main() {
	u1 := Guy{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := Guy{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := Guy{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []Guy{u1, u2, u3}

	fmt.Println(users)
	sort.Sort(ByAge(users))
	fmt.Println(users)
	sort.Sort(ByLast(users))
	fmt.Println(users)
}
