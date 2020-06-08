package main

import (
	"fmt"
)

func main() {
	a := (72 == 72)
	b := (82.2 <= 81.2)
	c := (42 >= 43)
	d := (72.6 != 72.0)
	e := (41 < 42)
	f := (42 < 41)
	g := ("ad" == "ad")
	fmt.Println(a, b, c, d, e, f, g)
}
