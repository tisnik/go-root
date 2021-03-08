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
