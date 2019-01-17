package main

import (
	"fmt"
)

func main() {
	f := func() {
		fmt.Println("My first func expression!")
	}
	f()

	f := func(x int) {
		fmt.Println("My Birthday: ")
	}
	f(1997)
}

/// func (r receiver) identifier(parameter(s)) (return(s)) { code}
