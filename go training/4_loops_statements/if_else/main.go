package main

import "fmt"

func main() {
	x := 13 % 3
	fmt.Println(x)
	if x == 1 {
		fmt.Println("odd")
	} else {
		fmt.Println("even")
	}
}
