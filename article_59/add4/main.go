package main

import "fmt"
import "add4/adder"

func main() {
	fmt.Println(adder.IntAdd(1, 2))
	fmt.Println(adder.Float32Add(1.1, 2.2))
}
