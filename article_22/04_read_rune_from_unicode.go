// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá druhá část
//     Vstupně-výstupní funkce standardní knihovny programovacího jazyka Go
//     https://www.root.cz/clanky/vstupne-vystupni-funkce-standardni-knihovny-programovaciho-jazyka-go/
//
// Demonstrační příklad číslo 4:
//     Načítání jednotlivých run (ne bajtů) z řetězce.

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("* ěščř ½µ§я¤ *")
	for {
		c, size, err := reader.ReadRune()
		if err == io.EOF {
			fmt.Println("\nEOF detected")
			break
		}
		if err == nil {
			fmt.Printf("%c %d\n", c, size)
		} else {
			fmt.Printf("\nerror %v", err)
			break
		}
	}
}
