// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá druhá část
//     Vstupně-výstupní funkce standardní knihovny programovacího jazyka Go
//     https://www.root.cz/clanky/vstupne-vystupni-funkce-standardni-knihovny-programovaciho-jazyka-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z dvacáté druhé části:
//    https://github.com/tisnik/go-root/blob/master/article_22/README.md
//
// Demonstrační příklad číslo 3:
//     Postupné načtení bajtů z řetězce, ukázka použití Unicode.

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("* ěščř ½µ§я¤ *")
	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			fmt.Println("\nEOF detected")
			break
		}
		if err == nil {
			fmt.Printf("%02x ", b)
		} else {
			fmt.Printf("\nerror %v", err)
			break
		}
	}
}
