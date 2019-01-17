package main

import "fmt"

func main() {
	i := 0
	for i < 10 { //i < 10 is the init
		fmt.Println(i)
		i++
	}
}

// for init; condition; post{}
// for init; while; make this happen{}
