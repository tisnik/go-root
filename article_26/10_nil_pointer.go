package main

import "fmt"

func main() {
	var v *int = nil

	fmt.Println(v)

	*v = 42
}
