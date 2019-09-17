package main

import "fmt"

func main() {
	x := 1

	for cond := true; cond; cond = x <= 10 {
		fmt.Printf("%d\n", x)
		x++
	}
}
