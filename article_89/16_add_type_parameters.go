package main

import "fmt"

type numeric interface {
	int | float64
}

func add[T numeric](x T, y T) T {
	var first T = x
	var second T = y
	return first + second
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(3.14, 6.28))
}
