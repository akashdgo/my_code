package main

import (
	"fmt"
)

func main() {
	x := 12
	fmt.Printf("%d\t\t\t%b\t\t\t%#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d\t\t\t%b\t\t\t%#x\n", y, y, y)
}
