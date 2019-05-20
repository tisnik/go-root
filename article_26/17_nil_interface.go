package main

import "fmt"

func main() {
	var i1 interface{}

	fmt.Printf("%v %v\n", i1, i1 == nil)
}
