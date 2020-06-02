package main

import (
	"fmt"
)

// declar + assign == initialize
var y = 77

// Declare a variable with the Identifier 'a'
// and the variable with the Identifier 'z' is of Type 'int'

// When storage is allocated for a variable, either through a declaration or a call of new,
// or when a new value is created, either through a composite literal or a call of make, and
// no explicit initialization is provided, the variable or value is given a default value.
// Each element of such a variable or value is set to the zero value for its type: false for
// booleans, 0 for numeric types, "" for strings, and nil for pointers, functions, interfaces,
// slices, channels, and maps.

var a int // Equivalent of 'var a int = 0'

func main() {
	// short declaration operator (since it operates on Operands)
	// declar a variable and assign a value (of a certain type)

	x := 42 // non-declaration statement and needs to be inside function body
	fmt.Println(x)
	fmt.Println("Print y before modification. y ==", y)

	// To declare a variable outside a function body use 'var'

	var z = y + 10
	y = y + 20
	fmt.Println("Print value of z ==", z)
	fmt.Println("New value y ==", y)
	foobar()
	fmt.Println(a)
}

func foobar() {
	//  y := 01  // short declaration operator just for inside function foobar()
	fmt.Println("Value of y used by foobar called from inside the body of main(): y ==", y)
}
