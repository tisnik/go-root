package main

import "fmt"

func main() {
	for x := 1; x <= 10000; x<<=1 {
		fmt.Printf("%d\n", x)
	}
}
