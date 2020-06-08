package main

import (
	"fmt"
)

type owntype int

var x owntype

func main() {
	fmt.Println(x)
	fmt.Printf("x is of TYPE %T\n", x)
	x = 42
	fmt.Println(x)
}
