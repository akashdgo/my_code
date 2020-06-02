package main

import (
	"fmt"
)

var x string
var y int

func main() {
	// DECLARE a variable to be of a certain TYPE
	// and then ASSIGN a VALUE of that particular type to the variable
	fmt.Println("Print the value of x", x, ".<-- a blank space")
	fmt.Printf("%T\n", x)

	x = "Akash Das"

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
