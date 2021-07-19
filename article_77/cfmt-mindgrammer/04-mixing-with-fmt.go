package main

import (
	"fmt"

	"github.com/mingrammer/cfmt"
)

func main() {
	fmt.Print("1st line: ")
	cfmt.Successln("Success message")

	fmt.Print("2nd line: ")
	cfmt.Infoln("Info message")

	fmt.Print("3rd line: ")
	cfmt.Warningln("Warning message")

	fmt.Print("4th line: ")
	cfmt.Errorln("Error message")

	fmt.Println()
	fmt.Println("That's all...")
}
