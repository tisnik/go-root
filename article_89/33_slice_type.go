package main

import "fmt"

type Slice[T any] []T

func car[T any](s Slice[T]) T {
	return s[0]
}

func cdr[T any](s Slice[T]) []T {
	return s[1:]
}

func (s Slice[T]) length() int {
	return len(s)
}

func main() {
	s := Slice[int]{1, 2, 3}

	fmt.Println(s)
	fmt.Println(car(s))
	fmt.Println(cdr(s))

	fmt.Println()

	fmt.Println(len(s))
	fmt.Println(len(cdr(s)))
}
