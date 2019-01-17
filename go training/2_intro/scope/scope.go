package main

import "fmt"

var x int = 1

//the variable x of type int = 1 applies to all of the package

func main() {

	y := 2

	fmt.Println(x)
	foo()
	fmt.Println(y)
	{
		v := 3
		fmt.Println(x)
		fmt.Println(y)
		fmt.Println(v)

		//a variable inside a scope within another scope will be ordered torwards the end of the output
	}
}

//the variable y := 2 must be above the command fmt.Println(y). Order matters in block level scope!

func foo() {
	fmt.Println(x)

}

//the x variable applies to func foo() because it applies to the whole file but the variable y will not because it isn't in the block level scope of func foo() or the file level scope
