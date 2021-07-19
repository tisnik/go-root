package main

import (
	"fmt"

	"github.com/mingrammer/cfmt"
)

func main() {
	var msg string

	msg = cfmt.Ssuccess("Success message")
	fmt.Printf("1st line: %s\n", msg)

	msg = cfmt.Sinfo("Info message")
	fmt.Printf("2nd line: %s\n", msg)

	msg = cfmt.Swarning("Warning message")
	fmt.Printf("3rd line: %s\n", msg)

	msg = cfmt.Serror("Error message")
	fmt.Printf("4th line: %s\n", msg)

	fmt.Println()
	fmt.Println("That's all...")
}
