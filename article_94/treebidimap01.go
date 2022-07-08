// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/

package main

import (
	"fmt"

	mapImpl "github.com/daichi-m/go18ds/maps/treebidimap"
	"github.com/daichi-m/go18ds/utils"
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

	m := mapImpl.NewWith(utils.StringComparator, utils.StringComparator)
	fmt.Println(m)

	for i := 0; i < len(cz); i++ {
		m.Put(cz[i], lat[i])
		fmt.Println(m)
	}
}
