package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	a := [3]int(s)
	a[0] = 9
	fmt.Println("slice:", s)
	fmt.Println("array: ", a)
}
