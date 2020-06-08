package main

import "fmt"

func main() {
	s := `vim
	-go`
	t := "string types"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	fmt.Println(t)

	// convert from type 'string' to 'slice byte'
	byte_str := []byte(t)
	fmt.Println(byte_str)
	fmt.Printf("%T\n", byte_str)

	for i := 0; i < len(t); i++ {
		// %#U gives the corresponding UTF-8 encoding for each 'index (i)' for var 't'
		fmt.Printf("%#U\n", t[i])
	}
	// for index position for the  range of 't' & then print it as hex (0x)
	for i, v := range t {
		fmt.Printf("The index of %d, the hex value is %#x\n", i, v)
	}
}
