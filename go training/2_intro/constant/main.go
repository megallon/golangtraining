package main

import "fmt"

const p string = "death and taxes"

const (
	Pi       = 3.14
	Language = "Go"
)

func main() {
	fmt.Println(Pi)
	fmt.Println(Language)
	fmt.Println("p -", p)
}

//a constant is a simple unchanging value
//There's a couple different ways to do constants like these two ways shown above
