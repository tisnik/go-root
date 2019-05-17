package main

import "fmt"

func main() {
	var s struct {
		a int
		b bool
		c chan int
		d []int
	}

	fmt.Println(s)
}
