package main

import "fmt"

func main() {
	greet("Jane")
	greet("John")
}

func greet(name string) {
	fmt.Println(name)
}

//func starts with func then recievers if used then the function name ex:main then () parameters then returns
//func-recievers-name-parameters-returns
//ex:greet is declared with a parameter, when calling greet pass in an argument
