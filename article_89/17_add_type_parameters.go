package main

import "fmt"

type numeric interface {
	int | float64
}

func add[T numeric](x T, y T) T {
	return x + y
}

func main() {
	fmt.Println(add(1, 3.14))
	fmt.Println(add(3.14, 1))
}
