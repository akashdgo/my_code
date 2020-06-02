package main

import (
	"fmt"
)

func main() {

	n, err := fmt.Println("Hello World", 111, true)
	fmt.Println(n)
	fmt.Println(err)

	// Use '_' to ignore a particular value being returned

	m, _ := fmt.Println("Hello Akash!!", 112, true)
	fmt.Println(m)
}
