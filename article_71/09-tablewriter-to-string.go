// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá první část
//    Formátovaný tisk obsahu tabulek a dalších datových struktur v Go
//    https://www.root.cz/clanky/formatovany-tisk-obsahu-tabulek-a-dalsich-datovych-struktur-v-go/

package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func main() {
	asString := &strings.Builder{}

	table, err := tablewriter.NewCSV(asString, "tiobe.csv", true)
	if err != nil {
		log.Fatal(err)
	}
	table.SetRowLine(true)
	table.SetAutoMergeCellsByColumnIndex([]int{1, 2})
	table.Render()

	fmt.Println(asString.String())
}
