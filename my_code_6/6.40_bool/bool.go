package main

import "fmt"

var x bool

func main() {
	fmt.Println("Print x when defined as 'var x bool'. x =", x)
	x = true
	fmt.Println("Print x when defined previous line as TRUE. x =", x)
	a := 7
	b := 42
	fmt.Println("a=7, b=42")
	fmt.Println("Print a == b ==>", a == b)
}
