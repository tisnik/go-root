package main

import "fmt"
import "add3/adder"

func main() {
	fmt.Println(adder.Add(1, 2))
	fmt.Println(adder.Add(1.1, 2.2))
}
