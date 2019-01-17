package main

import "fmt"

func main() {
	greet("Jane", "Doe")
}

func greet(fname, lname string) {
	fmt.Println(fname, lname)
}

//greet is declared with a parameter
//when calling greet you pass in an arguement
//func starts with func then recievers if used then the function name ex:main then () parameters then returns
//func-recievers-name-parameters-returns
//more functions allow you to abstract functionality
