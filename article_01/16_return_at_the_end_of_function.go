// Seriál "Programovací jazyk Go"
//
// Demonstrační příklad číslo 16:
//    Chybný pokus o deklaraci proměnné za příkazem return

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
