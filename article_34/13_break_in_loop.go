package main

import "fmt"

func main() {
	x := 1

	for { /* endless loop */
		fmt.Printf("%d\n", x)
		x++
		if x > 10 {
			break
		}
	}
}
