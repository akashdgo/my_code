package main

import (
	"fmt"
)

const (
	a int     = 42
	b float64 = 42.78
	c string  = "James Bond"
)

func main() {
	fmt.Println("Two types of consts:\n", "1. typed\n", "2. untyped")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
