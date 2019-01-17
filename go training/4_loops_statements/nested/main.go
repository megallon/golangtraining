package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i, "-", j)
		}
	}
}

//The j loop will complete in order to trigger the i loop
