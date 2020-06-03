package main

import (
	"fmt"
	"io"
	"os"
)

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#X\n", y)
	fmt.Printf("%o\n", y)
	fmt.Printf("%O\n", y)
	y = 911
	fmt.Printf("%#x\n", y)
	fmt.Printf("%x\t%#x\t%#X\t%b\n", y, y, y, y)

	fmt.Printf("%#x\t%b\t%x\t%#o\t%O\t%#X\n", y, y, y, y, y, y)

	s := fmt.Sprintf("%#x\t%b\t%x\t%#o\t%O\t%#X\n", y, y, y, y, y, y)
	fmt.Println(s)
	fmt.Printf("%T\n\n", s)

	fmt.Println("-------------------------------------------------------")
	const name, age = "Akash", 43
	fmt.Println("func Print")
	fmt.Print(name, " is ", age, " years old.\n")

	// It is conventional not to worry about any
	// error returned by Print.

	fmt.Println("-------------------------------------------------------")
	fmt.Println("func Printf")
	fmt.Printf("%s is %d years old.\n", name, age)

	// It is conventional not to worry about any
	// error returned by Printf.

	fmt.Println("-------------------------------------------------------")
	fmt.Println("func Println")
	fmt.Println(name, "is", age, "years old.")

	// It is conventional not to worry about any
	// error returned by Println.

	fmt.Println("-------------------------------------------------------")
	fmt.Println("func Sprint")
	v := fmt.Sprint(name, " is ", age, " years old.\n")
	io.WriteString(os.Stdout, v) // Ignoring error for simplicity.

	fmt.Println("-------------------------------------------------------")
	fmt.Println("func Sprintf")
	t := fmt.Sprintf("%s is %d years old.\n", name, age)
	io.WriteString(os.Stdout, t) // Ignoring error for simplicity.

	fmt.Println("-------------------------------------------------------")
	fmt.Println("func Sprintln")
	u := fmt.Sprintln(name, "is", age, "years old.")
	io.WriteString(os.Stdout, u) // Ignoring error for simplicity.

}
