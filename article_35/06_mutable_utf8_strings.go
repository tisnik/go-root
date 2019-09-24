package main

import "fmt"

func main() {
	var s []byte = []byte("[шщэюя]")

	fmt.Println(string(s))

	s[0] = '*'

	// problematicka cast
	s[5] = '-'

	s[11] = '*'
	fmt.Println(string(s))
}
