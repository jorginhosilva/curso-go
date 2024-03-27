package main

import (
	"fmt"
) 

type asdf int
var x asdf

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}