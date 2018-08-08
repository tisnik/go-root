package main

import "fmt"

func getMessage() (message string) {
	message = "Hello world!"
	return
}

func printMessage(message string) {
	fmt.Println(message)
}

func main() {
	printMessage(getMessage())
}
