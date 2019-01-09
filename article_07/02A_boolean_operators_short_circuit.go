// Seriál "Programovací jazyk Go"
//
// Sedmá část
//
// Demonstrační příklad číslo 2:
//    Logické operátory.

package main

import "fmt"

func f1() bool {
	println("f1")
	return true
}

func f2() bool {
	println("f2")
	return false
}

func f3() bool {
	println("f2")
	return false
}

func main() {
	fmt.Printf("short circuit &&: %v\n", f1() && f2())
	fmt.Printf("short circuit ||: %v\n", f1() || f2())
	fmt.Printf("short circuit &&: %v\n", f2() && f3())
	fmt.Printf("short circuit ||: %v\n", f2() || f3())
}
