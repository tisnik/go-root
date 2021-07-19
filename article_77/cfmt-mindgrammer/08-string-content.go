package main

import (
	"encoding/hex"
	"fmt"

	"github.com/mingrammer/cfmt"
)

func printEncoded(s string) {
	bytes := []byte(s)
	fmt.Printf("Encoded:\n%s\n", hex.Dump(bytes))
}

func main() {
	var msg string

	msg = cfmt.Ssuccessln("Success message")
	fmt.Printf("1st line: %s", msg)
	printEncoded(msg)

	msg = cfmt.Sinfoln("Info message")
	fmt.Printf("2nd line: %s", msg)
	printEncoded(msg)

	msg = cfmt.Swarningln("Warning message")
	fmt.Printf("3rd line: %s", msg)
	printEncoded(msg)

	msg = cfmt.Serrorln("Error message")
	fmt.Printf("4th line: %s", msg)
	printEncoded(msg)

	fmt.Println()
	fmt.Println("That's all...")
}
