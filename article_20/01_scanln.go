// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá část
//     Jazyk Go a aplikace s vlastním příkazovým řádkem
//     https://www.root.cz/clanky/jazyk-go-a-aplikace-s-vlastnim-prikazovym-radkem/
//
// Seznam příkladů z dvacáté části:
//    https://github.com/tisnik/go-root/blob/master/article_20/README.md
//
// Demonstrační příklad číslo 1:
//     Načtení řetězců ze standardního vstupu.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_20/01_scanln.html
//

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
