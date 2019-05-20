package main

import "fmt"

func main() {
	var p1 *int
	var p2 *int

	fmt.Printf("%v %v\n", p1, p1 == nil)
	fmt.Printf("%v %v\n", p2, p2 == nil)
	fmt.Printf("%v\n", p1 == p2)
}
