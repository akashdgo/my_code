package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello world", 111, true)
	fmt.Println(n)
	fmt.Println(err)
}
