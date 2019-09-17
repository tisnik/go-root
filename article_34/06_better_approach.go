package main

import "fmt"

func main() {
	x := 10

	switch {
	case x > 0:
		fmt.Println("x is positive number")
	case x == 0:
		fmt.Println("x is zero")
	default:
		fmt.Println("x is negative number")
	}
}
