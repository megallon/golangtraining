package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secAgent struct {
	person
	ltk bool
}

func (s secAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func main() {
	sa1 := secAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa2 := secAgent{
		person: person{
			"James",
			"Boss",
		},
		ltk: false,
	}
	sa1.speak()
	sa2.speak()
}

/// func (r receiver) identifier(parameter(s)) (return(s)) { code}
