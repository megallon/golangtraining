package main

import "fmt"

var g string = "go" //This is an example of a declaration and assignment -
//Declaring variable g is of type string and then assigning it to "go"

func main() {
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", g)
	// %T respresentation of the type of the value

}
