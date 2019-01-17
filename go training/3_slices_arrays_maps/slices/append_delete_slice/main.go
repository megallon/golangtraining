package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42} // composite literal
	fmt.Println(x)
	x = append(x, 77, 88, 99, 114, 1014) // append returns a slice of that type
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x) // deleting 7 and 8 from a slice
}

//... before a variable means it takes an unlimited number values of this type of this function
//... after means take the values from something and put them after your variable you ... after
