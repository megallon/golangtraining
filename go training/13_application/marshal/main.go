package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

// if uppercased different variables will be exported

func main() {
	s := `[{"First":"Austin","Last":"Megarroni","Age":21},{"First":"Morgan","Last":"Beckerroni","Age":19}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	p1 := person{
		First: "Austin",
		Last:  "Megarroni",
		Age:   21,
	}

	p2 := person{
		First: "Morgan",
		Last:  "Beckerroni",
		Age:   19,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
