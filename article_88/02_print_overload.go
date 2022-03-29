package main

import "fmt"

func printValue(value string) {
	fmt.Println(value)
}

func printValue(value int) {
	fmt.Println(value)
}

func main() {
	printValue("www.root.cz")
}
