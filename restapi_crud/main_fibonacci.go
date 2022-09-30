package main

import (
	"fmt"
)

func main() {
	var end int
	fmt.Print("Input max fibonacci number to print: ")
	_, err := fmt.Scanf("%d", &end)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		printFibonacci(end)
	}
}

func printFibonacci(end int) {
	var first int = 0
	var second int = 1
	var third int = first + second

	fmt.Print(first, "\t", second)

	for third < end {
		fmt.Print("\t", third)
		first = second
		second = third
		third = first + second
	}
}
