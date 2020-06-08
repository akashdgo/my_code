package main

import (
	"fmt"
)

const (
	_  = iota
	y1 = iota + 2016
	y2 = iota + 2016
	y3 = iota + 2016
	y4 = iota + 2016
)

func main() {
	fmt.Println(y1)
	fmt.Println(y2)
	fmt.Println(y3)
	fmt.Println(y4)
}
