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
// Demonstrační příklad číslo 6:
//     Zápis bajtů du bufferu, přečtení ve formě run.

package main

import (
	"bytes"
	"fmt"
)

func main() {
	buffer := bytes.Buffer{}
	buffer.WriteByte(65)

	c, size, err := buffer.ReadRune()
	fmt.Printf("%c %d %v\n", c, size, err)

	buffer.WriteByte(0xc4)
	buffer.WriteByte(0x9b)

	c, size, err = buffer.ReadRune()
	fmt.Printf("%c %d %v\n", c, size, err)

	c, size, err = buffer.ReadRune()
	fmt.Printf("%c %d %v\n", c, size, err)
}
