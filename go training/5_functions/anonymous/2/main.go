package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x

	}
}

//The second func is an anonymous func
//Returning a function

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}

//changing x will be remembered by outer scopes when called on like it is in increment := wrapper() that's why from func wrapper() func() int{} to func main() x works
