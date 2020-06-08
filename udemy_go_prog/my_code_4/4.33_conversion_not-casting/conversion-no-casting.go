package main

import (
	"fmt"
)

var x int

type owntype int

var y owntype

func main() {
	x = 77
	y = 81
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	// x =! y coz x is type'int' and y is type 'owntype'
	x = int(y)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
