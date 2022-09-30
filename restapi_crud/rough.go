package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	var myint int8 = 2
	for i := 0; i < 129; i++ {
		myint += 1
	}
	fmt.Println(myint) // this oddly prints out -125!
}
