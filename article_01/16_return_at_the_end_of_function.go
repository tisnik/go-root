package main

import "fmt"

func getMessage() (message string) {
	message = "Hello world!"
	return
	x := 42
}

func printMessage(message string) {
	fmt.Println(message)
}

func main() {
	printMessage(getMessage())
}
