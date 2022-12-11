package main

import "fmt"

func changeMe(slice []int) {
	slice[0] = 42
}

func main() {
	s := []int{1, 2, 3}
	fmt.Println(s)
	changeMe(s)
	fmt.Println(s)
}
