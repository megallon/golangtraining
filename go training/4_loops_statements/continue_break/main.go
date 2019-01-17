package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		} //if i is divided by 2 and has remainder of 0 continue; if not go back and do the next loop
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}
}
