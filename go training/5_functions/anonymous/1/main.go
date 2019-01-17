package main

import (
	"fmt"
)

func main() {
	func(x int) {
		fmt.Println("The meaning of life: ", x)
	}(42)

}

/// func (r receiver) identifier(parameter(s)) (return(s)) { code}
