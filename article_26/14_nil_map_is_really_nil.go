package main

import "fmt"

func main() {
	var m1 map[string]int = nil
	var m2 map[string]int

	fmt.Printf("%v %v\n", m1, m1 == nil)
	fmt.Printf("%v %v\n", m2, m2 == nil)
}
