// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá druhá část
//    Vstupně-výstupní funkce standardní knihovny programovacího jazyka Go
//    https://www.root.cz/clanky/vstupne-vystupni-funkce-standardni-knihovny-programovaciho-jazyka-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z dvacáté druhé části:
//    https://github.com/tisnik/go-root/blob/master/article_22/README.md
//
// Demonstrační příklad číslo 5:
//    Struktura splňující rozhraní Writer.
//    Zápis bajtu do bufferu.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_22/03_read_byte_from_unicode.html

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
