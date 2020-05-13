// Seriál "Programovací jazyk Go"
//
// Dvacátá část
//     Jazyk Go a aplikace s vlastním příkazovým řádkem
//     https://www.root.cz/clanky/jazyk-go-a-aplikace-s-vlastnim-prikazovym-radkem/
//
// Demonstrační příklad číslo 2:
//     Použití metody Reader.ReadString pro načítání ze standardního vstupu.

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
