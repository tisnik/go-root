package main

import "fmt"

func compare(x int, y int) bool {
	return x < y
}

func main() {
	fmt.Println(compare(1, 2))
}
