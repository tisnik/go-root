package main

import "fmt"

func main() {
	s := "[шщэюя]"

	var r []rune = []rune(s)

	fmt.Println(string(r))

	r[0] = '*'

	r[3] = '-'

	r[6] = '*'
	fmt.Println(string(r))
}
