package main

import "fmt"

func main() {
	var v *int = 1234

	fmt.Println(v)

	*v = 42
}
