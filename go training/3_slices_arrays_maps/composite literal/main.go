package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42} // composite literal
	fmt.Println(len(x))
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	for i, v := range x {
		fmt.Println(i, v)
	}
	//for range can be used to get the index and the value represented here as i, v
	//fmt.Println(x[0]) is an example of accessing from index position. (x[0]) accesses 4
	//x := type{values}
	//a slice allows you to group together values of the same type
}
