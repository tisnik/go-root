// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá část
//     Jazyk Go a aplikace s vlastním příkazovým řádkem
//     https://www.root.cz/clanky/jazyk-go-a-aplikace-s-vlastnim-prikazovym-radkem/
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
