package main

import "fmt"

func main() {
	s := "шщэюя"

	fmt.Println(s[2:4])
	fmt.Println(s[2:5])
	fmt.Println(s[2:6])

	fmt.Println()
	fmt.Println(s[3:4])
	fmt.Println(s[3:5])
	fmt.Println(s[3:6])
}
