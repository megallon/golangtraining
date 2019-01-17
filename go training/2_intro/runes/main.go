package main

import "fmt"

func main() {
	switch "Marcus" {
	//looks for the switch when going through cases
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Marcus":
		fmt.Println("Wassup Marcus")
		fallthrough
		//falls through and enters the next line of code
	case "Jack":
		fmt.Println("Wassup Jack")
		fallthrough
	case "Julian":
		fmt.Println("Wassup Julian")
	case "Bob":
		fmt.Println("Wassup Bob")
	default:
		fmt.Println("No friends?")
	}
}
