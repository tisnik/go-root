package main

import "fmt"

func main() {
	var name string

	fmt.Print("Name: ")
	fmt.Scanln(&name)

	fmt.Printf("Hello, %s!\n", name)
}
