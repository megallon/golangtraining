package main

import "fmt"

var g string = "go"//This is an example of a declaration and assignment -
//Declaring variable g is of type string and then assigning it to "go"

func main() {
	a := 10 //These are examples of shorthand variables
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", g)
}
