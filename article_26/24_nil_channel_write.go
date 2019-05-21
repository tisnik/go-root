package main

import "fmt"

func main() {
	var c1 chan int = nil

	fmt.Printf("%v %v\n", c1, c1 == nil)

	c1 <- 10
}
