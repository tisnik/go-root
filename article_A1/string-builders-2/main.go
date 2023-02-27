package main

import (
	"fmt"
)

func main() {
	const n = 20

	s := BuildStringUsingConcatenation(n)
	fmt.Println("Concatenation:                   ", s)

	s2 := BuildStringUsingStringBuffer(n)
	fmt.Println("StringBuffer (initially empty):  ", s2)

	s3 := BuildStringUsingPreallocatedStringBuffer(n)
	fmt.Println("StringBuffer (preallocated):     ", s3)

	s4 := BuildStringUsingStringBuilder(n)
	fmt.Println("StringBuilder (initially empty): ", s4)

	s5 := BuildStringUsingStringBuilder(n)
	fmt.Println("StringBuilder (preallocated):    ", s5)
}
