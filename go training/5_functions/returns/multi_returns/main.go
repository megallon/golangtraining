package main

import "fmt"

func main() {
	fmt.Println(greet("Jane", "Doe"))
}

func greet(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}

//fmt.Sprint is string print. Brings arguments together and returns a string.
//func starts with func then recievers if used then the function name ex:main then () parameters then returns
//func-recievers-name-parameters-returns
//ex:greet is declared with a parameter, when calling greet pass in an argument
