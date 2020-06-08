package main

import "fmt"

func main() {
	s := "H"
	fmt.Println("value of 's' =", s)

	// byte-slice of 'H'
	byte_slice := []byte(s)
	fmt.Println("byte slice value of 'H'= ", byte_slice)

	//variable for 0th position in byte_slice
	n := byte_slice[0]
	fmt.Println("value of first value in byte slice of 's', n = ", n)
	fmt.Printf("type of 'n'= %T\n", n)
	fmt.Printf("type of 'n' in binary = %b\n", n)
	fmt.Printf("type of 'n' in oct = %O\n", n)
	fmt.Printf("type of 'n' in hex = %#x\n", n)
}
