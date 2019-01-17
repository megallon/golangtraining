package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 42 //change the value at this memory address to 42

	fmt.Println(a) //a will now print 42
}

//b is an int pointer
//b is a pointer to the memory address where an int is stored
//b is of type "int pointer"
//* is part of the type - b is of type *int
//to see the value in that memory address where an int is stored put a * in front of b to dereference the pointer and see the value
