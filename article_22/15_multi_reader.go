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
// Demonstrační příklad číslo 15:
//     Multi reader se třemi vstupy.

package main

import (
	"fmt"
	"io"
	"strings"
)

const inputString1 = "Hello"
const inputString2 = "world"
const inputString3 = "!"

const bufferSize = 4

func main() {
	r1 := strings.NewReader(inputString1)
	r2 := strings.NewReader(inputString2)
	r3 := strings.NewReader(inputString3)
	r := io.MultiReader(r1, r2, r3)

	buffer := make([]byte, bufferSize)

	for {
		read, err := r.Read(buffer)

		if read > 0 {
			fmt.Printf("read %d bytes\n", read)
			fmt.Println(string(buffer[:read]))
		}

		if err == io.EOF {
			fmt.Println("reached end of file")
			break
		}

		if err != nil {
			fmt.Printf("other error %v\n", err)
			break
		}
	}
}
