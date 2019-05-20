package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	fmt.Printf("%v %v\n", m1, m1 == nil)

	m1["foo"] = 3
	fmt.Printf("%v %v\n", m1, m1 == nil)
}
