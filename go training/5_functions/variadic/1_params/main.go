package main

import "fmt"

func main() {
	data := []float64{2, 4}//[] makes this into a slice of bites instead of individual kinds
	n := average(data...)
	fmt.Println(n)
}

func average(sf ...float64) float64 { 
	fmt.Println(sf)
	fmt.Printf("%T", sf)
	var total float64 // total is an accumulator, it accumulates values as you go through a loop
	for _, v := range sf { // range allow you to range over sf
		return total += v // total += v is the same as total = total + v
	}
	return total / float64 (len(sf))
}

	