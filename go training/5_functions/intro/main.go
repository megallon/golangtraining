package main

import "fmt"

func main() {
bar("James")
s1 := woo("Money")
fmt.Println(st)
x, y := mouse("Ian", "Flemming")
}

func bar(s string) {
	fmt.Println("Hello", s)
}

func woo(st string){
	return fmt.Sprint("Hello from woo, ", st)
}

func mouse(fn string, ln string)(string, bool) {
	a := fmt.Sprint(fn, ln ", says Hello")
	b := true
	return a, b

}

//greet is declared with a parameter
//when calling greet you pass in an arguement
//func starts with func then recievers if used then the function name ex:main then () parameters then returns
//func(recievers)identifier(parameters)(return(s)) {...}
//more functions allow you to abstract functionality
