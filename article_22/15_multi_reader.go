// Seriál "Programovací jazyk Go"
//
// Dvacátá druhá část
//     Vstupně-výstupní funkce standardní knihovny programovacího jazyka Go
//     https://www.root.cz/clanky/vstupne-vystupni-funkce-standardni-knihovny-programovaciho-jazyka-go/
//
// Demonstrační příklad číslo 15:
//     Multi reader se třemi vstupy.

package main

import (
	"fmt"
	"io"
	"strings"
)

const input_string_1 = "Hello"
const input_string_2 = "world"
const input_string_3 = "!"

const bufferSize = 4

func main() {
	r1 := strings.NewReader(input_string_1)
	r2 := strings.NewReader(input_string_2)
	r3 := strings.NewReader(input_string_3)
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
