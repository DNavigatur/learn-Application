package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Last  string
	Age   int
}

/*
func main() {

	p1 := Person{
		First: "James",
		Last:  "Bond",
		Age:   64,
	}
	p2 := Person{
		First: "Dwayne",
		Last:  "Johnson",
		Age:   49,
	}

	People := []Person{p1, p2}

	fmt.Println(People)

	bs, err := json.Marshal(People)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}

// hands on exercise:- Marshal to json format

type User struct {
	First string
	Age   int
}

func main() {
	u1 := User{
		First: "James",
		Age:   32,
	}

	u2 := User{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := User{
		First: "M",
		Age:   54,
	}

	Users := []User{u1, u2, u3}

	fmt.Println(Users)

	// your code goes here
	bs, err := json.Marshal(Users)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	fmt.Printf("%T", bs)
}

// hands on exercise:- Unmarshal to json format

type Persun struct {
	First   string   `json:"First,omitempty"`
	Last    string   `json:"Last,omitempty"`
	Age     int      `json:"Age,omitempty"`
	Sayings []string `json:"Sayings,omitempty"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	//fmt.Println(s)

	bs := []byte(s)

	var People []Persun

	err := json.Unmarshal(bs, &People)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(People)

	for _, peep := range People {
		fmt.Println("\t", peep.First, peep.Last, peep.Age)
		for _, say := range peep.Sayings {
			fmt.Println("\t\t", say)
		}
	}
}

// hands on exercise:- encode to json format

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	// json package is used to encode the users slice as JSON and write the JSON data to the standard output. The json.NewEncoder function creates a new JSON encoder, and os.Stdout is used as the output destination (standard output). The .Encode(users) call encodes the users slice and writes it to the standard output.
	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println(err)
	}
}

// hands on exercise:- sort to json format

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs
	sort.Strings(xs)
	fmt.Println(xs)
}
*/
// hands on exercise:- sort costum to json format

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// interface logic for ByAge sorting
type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// interface logic for ByLast sorting
type ByLast []user

func (l ByLast) Len() int           { return len(l) }
func (l ByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }
func (l ByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	//fmt.Println(users)

	// your code goes here
	//sort ByAge
	sort.Sort(ByAge(users))
	fmt.Println("==============================")
	//fmt.Println(users)
	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		for _, s := range v.Sayings {
			fmt.Println("\n", s)
		}
	}

	//sort ByLast
	sort.Sort(ByLast(users))
	fmt.Println("==============================")
	//fmt.Println(users)
	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		for _, s := range v.Sayings {
			fmt.Println("\n", s)
		}
	}

	//sort by sayings
	fmt.Println("==============================")
	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		sort.Strings(v.Sayings)
		for _, s := range v.Sayings {
			fmt.Println("\n", s)
		}
	}
}
