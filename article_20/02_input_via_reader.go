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
// Demonstrační příklad číslo 2:
//     Použití metody Reader.ReadString pro načítání ze standardního vstupu.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_20/02_input_via_reader.html
//

package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	print("Login: ")
	login, err := reader.ReadString('\n')
	if err != nil {
		println("Error reading login")
	}

	print("Password: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		println("Error reading password")
	}

	println(login)
	println(password)
}
