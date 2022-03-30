package main

import "fmt"

type numeric interface {
	int | float64
}

func add[T numeric, U numeric](x T, y U) T {
	return x + T(y)
}

func main() {
	fmt.Println(add(1, 3.14))
	fmt.Println(add(3.14, 1))
}
