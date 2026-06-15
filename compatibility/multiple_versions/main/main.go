package main

import (
	"fmt"

	"example.com/module1"
	"example.com/module2"
)

func main() {
	module1.Foo()
	fmt.Println()
	module2.Foo()
	fmt.Println()
}
