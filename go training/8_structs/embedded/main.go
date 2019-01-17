package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type strongboi struct {
	person
	lift bool
}

func main() {
	sb := strongboi{
		person: person{
			first: "Austin",
			last:  "Megallon",
			age:   21,
		},
		lift: true,
	}
	p2 := person{
		first: "Morgan",
		last:  "Beck",
		age:   19,
	}
	fmt.Println(sb)
	fmt.Println(p2)
	fmt.Println(sb.first, sb.last, sb.age, sb.lift)
	fmt.Println(p2.first, p2.last, p2.age)
}

//embedding one struct into another struct
