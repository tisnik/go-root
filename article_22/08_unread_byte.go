// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá druhá část
//     Vstupně-výstupní funkce standardní knihovny programovacího jazyka Go
//     https://www.root.cz/clanky/vstupne-vystupni-funkce-standardni-knihovny-programovaciho-jazyka-go/
//
// Demonstrační příklad číslo 8:
//     Použití metody UnreadByte.

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
		if cnt == 5 || cnt == 10 || cnt == 14 || cnt == 15 {
			reader.UnreadByte()
		}
	}
}
