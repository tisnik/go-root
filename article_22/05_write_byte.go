// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá druhá část
//     Vstupně-výstupní funkce standardní knihovny programovacího jazyka Go
//     https://www.root.cz/clanky/vstupne-vystupni-funkce-standardni-knihovny-programovaciho-jazyka-go/
//
// Seznam příkladů z dvacáté druhé části:
//    https://github.com/tisnik/go-root/blob/master/article_22/README.md
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
