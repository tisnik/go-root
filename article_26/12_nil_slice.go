package main

import "fmt"

func main() {
	var s0 []int
	s1 := []int{}
	s2 := make([]int, 0)
	s3 := make([]int, 10)

	fmt.Println(s0)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
