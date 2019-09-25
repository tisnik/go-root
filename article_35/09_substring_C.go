package main

import "fmt"

func main() {
	s := "шщэюя"

	runes := []rune(s)

	fmt.Println(string(runes[2:4]))
	fmt.Println(string(runes[2:5]))
	fmt.Println(string(runes[2:6]))

	fmt.Println()
	fmt.Println(string(runes[3:4]))
	fmt.Println(string(runes[3:5]))
	fmt.Println(string(runes[3:6]))
}
