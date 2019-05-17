package main

import "fmt"

func main() {
	fmt.Println(true)
	fmt.Println(false)

	x := true
	true := false
	false := x

	fmt.Println("oh no...")

	fmt.Println(true)
	fmt.Println(false)
}
