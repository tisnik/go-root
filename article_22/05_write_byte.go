// Seriál "Programovací jazyk Go"
//
// Dvacátá druhá část
//     Vstupně-výstupní funkce standardní knihovny programovacího jazyka Go
//     https://www.root.cz/clanky/vstupne-vystupni-funkce-standardni-knihovny-programovaciho-jazyka-go/
//
// Demonstrační příklad číslo 5:
//     Zápis bajtu do bufferu.

package main

import (
	"bytes"
	"fmt"
)

func main() {
	buffer := bytes.Buffer{}
	buffer.WriteByte(65)

	b, _ := buffer.ReadByte()
	fmt.Printf("%02x\n", b)

	buffer.WriteByte(0xff)

	b, _ = buffer.ReadByte()
	fmt.Printf("%02x\n", b)
}
