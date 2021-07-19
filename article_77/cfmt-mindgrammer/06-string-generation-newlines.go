package main

import (
	"fmt"

	"github.com/mingrammer/cfmt"
)

func main() {
	var msg string

	msg = cfmt.Ssuccessln("Success message")
	fmt.Printf("1st line: %s", msg)

	msg = cfmt.Sinfoln("Info message")
	fmt.Printf("2nd line: %s", msg)

	msg = cfmt.Swarningln("Warning message")
	fmt.Printf("3rd line: %s", msg)

	msg = cfmt.Serrorln("Error message")
	fmt.Printf("4th line: %s", msg)

	fmt.Println()
	fmt.Println("That's all...")
}
