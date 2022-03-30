package main

import "fmt"

func car[T any](s []T) T {
	return s[0]
}

func cdr[T any](s []T) []T {
	return s[1:]
}

func main() {
	s := []int{1, 2, 3}
	fmt.Println(s)
	fmt.Println(car(s))
	fmt.Println(cdr(s))
}
