package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	fmt.Println(x, y, z)
	fmt.Printf("x is of type %T and zero value of x = %d\n", x, x)
	fmt.Printf("y is of type %T and zero value of y = %s\n", y, y)
	fmt.Printf("z is of type %T and zero value of z = %t\n", z, z)
	fmt.Printf("z is of type %T and zero value of z = %#v\n", z, z)
}
