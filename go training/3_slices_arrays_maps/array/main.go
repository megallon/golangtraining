package main

import "fmt"

func main() {
	var x [5]int
	var y [6]int //x and y are of two different types because length is part of an array's type
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(y)

}
