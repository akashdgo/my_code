package main

import (
	"fmt"
)

type owntype int

var x owntype
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("x is of TYPE '%T'\n", x)
	fmt.Printf("y is of TYPE '%T'\n", y)
	x = 42
	fmt.Println("x =", x)
	fmt.Println("---------------------------------------")
	fmt.Println("Conversion of assigning value of x to y")
	fmt.Println("---------------------------------------")
	y = int(x)
	fmt.Println("y =", y)
	//fmt.Printf("y is of TYPE '%T'\n", y)
}
