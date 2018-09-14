package main

import (
	"fmt"
	"strconv"
)

var global = 16

func main() {
	// var i = 19 // if i is not "used", the program will not compile
	x := 2     // x has type int
	// c := float32(x) + 1.1 // does not compile because of adding float to an int
	fmt.Printf("%T\n", x)

	// update global variable
	fmt.Printf("global = %v\n", global)
	global = 22
	fmt.Printf("global = %v\n", global)

	// string operation
	s := "GDG is #"
	// s = s + 1               // no implicit string concatenation!
	s = s + strconv.Itoa(1) // or
	s = fmt.Sprintf("GDG is #%d", 1)
	fmt.Println(s)
}
