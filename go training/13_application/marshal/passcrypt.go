package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//had to preform go get golang.org/x/crypto/bcrypt to import this package

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)
}
