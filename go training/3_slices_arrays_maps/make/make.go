package main

import "fmt"

func main() {
	x := make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 42
	x[9] = 999
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

//make([]Type, length, capacity)
//... before a variable means it takes an unlimited number values of this type of this function
//... after means take the values from something and put them after your variable you ... after
