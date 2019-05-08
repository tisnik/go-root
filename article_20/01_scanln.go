// Seriál "Programovací jazyk Go"
//
// Dvacátá část část
//
// Demonstrační příklad číslo 1:
//     Načtení řetězců ze standardního vstupu.

package main

import "fmt"

func main() {
	var login string
	var password string

	print("Login: ")
	n, err := fmt.Scanln(&login)
	if n != 1 || err != nil {
		println("Error reading login")
	}

	print("Password: ")
	n, err = fmt.Scanln(&password)
	if n != 1 || err != nil {
		println("Error reading password")
	}

	println(login)
	println(password)
}
