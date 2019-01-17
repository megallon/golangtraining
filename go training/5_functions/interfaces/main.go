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
	fmt.Println("I am", s.first, s.last, " - the secAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", s.first, s.last, "- the person speak")
}

type human interface{
	speak()
}
func bar(h human) {
	swtich h.(type) {
	case person:
		fmt.Println("I was passed into barrrrr", h.(person).first)
	case secAgent:fmt.Println("I was passed into barrrrr", h.(secAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}

type hotdog int{

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
	p1 := person{
		first: "Dr.",
		last: "Yes",
	}
	sa1.speak()
	sa2.speak()
	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)

	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)


}

/// func (r receiver) identifier(parameter(s)) (return(s)) { code}
// interfaces allow
