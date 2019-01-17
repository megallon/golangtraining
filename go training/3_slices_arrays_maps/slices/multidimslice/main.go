package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}

//make([]Type, length, capacity)
//... before a variable means it takes an unlimited number values of this type of this function
//... after means take the values from something and put them after your variable you ... after
