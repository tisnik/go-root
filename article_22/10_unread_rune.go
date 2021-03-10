// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá druhá část
//     Vstupně-výstupní funkce standardní knihovny programovacího jazyka Go
//     https://www.root.cz/clanky/vstupne-vystupni-funkce-standardni-knihovny-programovaciho-jazyka-go/
//
// Demonstrační příklad číslo 10:
//     Použití metody UnreadRune.

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("* ěščř ½µ§я¤ *")
	cnt := 0

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
		cnt++
		if cnt == 5 || cnt == 10 || cnt == 14 || cnt == 15 {
			reader.UnreadRune()
		}
	}
}
