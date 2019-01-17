package main

import "fmt"

func main() {
	x := 13 % 3 //13 / 3 would give 13 divided by 3;
	//13 % 3 gives the remainder
	fmt.Println(x)
	if x == 1 { //== checks to see if x is equal to 1
		fmt.Println("odd")
	} else { //if x is something "else" it prints even
		fmt.Println("even")
	}
}
