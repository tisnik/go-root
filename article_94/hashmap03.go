// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devadesátá čtvrtá část
//    Knihovny s implementací generických datových typů pro programovací jazyk Go (3)
//    https://www.root.cz/clanky/knihovny-s-implementaci-generickych-datovych-typu-pro-programovaci-jazyk-go-3/
//
// Seznam příkladů ze devadesáté čtvrté části:
//    https://github.com/tisnik/go-root/blob/master/article_94/README.md

package main

import (
	"fmt"

	mapImpl "github.com/daichi-m/go18ds/maps/hashmap"
)

func main() {
	cz := []string{
		"Leden",
		"Únor",
		"Březen",
		"Duben",
		"Květen",
		"Červen",
		"Červenec",
		"Srpen",
		"Září",
		"Říjen",
		"Listopad",
		"Prosinec",
	}

	lat := []string{
		"Ianuarius",
		"Februarius",
		"Martius",
		"Aprilis",
		"Maius",
		"Iunius",
		"Iulius",
		"Augustus",
		"September",
		"October",
		"November",
		"December",
	}

	m := mapImpl.New[string, string]()

	for i := 0; i < len(cz); i++ {
		m.Put(cz[i], lat[i])
	}

	fmt.Println("Metadata")
	fmt.Println("Empty?", m.Empty())
	fmt.Println("Size", m.Size())

	fmt.Println()
	fmt.Println("Keys")
	fmt.Printf("%T\n", m.Keys())
	fmt.Println(m.Keys())

	fmt.Println()
	fmt.Println("Values")
	fmt.Printf("%T\n", m.Values())
	fmt.Println(m.Values())
}
