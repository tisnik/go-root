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
// Demonstrační příklad číslo 2:
//    Struktura splňující rozhraní Reader.
//    Postupné načtení bajtů z řetězce, lepší reakce na chyby.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_22/02_better_read_byte.html

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello world!")
	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			fmt.Println("\nEOF detected")
			break
		}
		if err == nil {
			fmt.Printf("%c", b)
		} else {
			fmt.Printf("\nerror %v", err)
			break
		}
	}
}
