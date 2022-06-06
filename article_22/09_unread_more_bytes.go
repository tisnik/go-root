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
// Demonstrační příklad číslo 9:
//     Opakované použití metody UnreadByte.

package main

import (
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello world!")
	cnt := 0

	for {
		b, err := reader.ReadByte()
		if err == nil {
			fmt.Printf("%c", b)
		} else {
			fmt.Printf("\nerror %v", err)
			break
		}
		cnt++
		if cnt == 6 {
			for i := 0; i <= 6; i++ {
				reader.UnreadByte()
			}
		}
	}
}
