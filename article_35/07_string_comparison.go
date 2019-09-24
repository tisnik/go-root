package main

import "fmt"

func main() {
	fmt.Println("aa" < "ab")
	fmt.Println("aa" == "ab")

	fmt.Println("aa" < "aa")
	fmt.Println("aa" == "aa")

	fmt.Println("e" < "é")
	fmt.Println("e" < "ě")

	fmt.Println("z" < "é")
	fmt.Println("z" < "ě")

	fmt.Println("h" < "ch")
	fmt.Println("ch" < "i")

	fmt.Println("Hrdina" < "Chocholoušek")
	fmt.Println("Crha" < "Chocholoušek")
}
