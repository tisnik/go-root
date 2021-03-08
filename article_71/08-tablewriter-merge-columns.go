package main

import (
	"fmt"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	table, err := tablewriter.NewCSV(os.Stdout, "tiobe.csv", true)
	if err != nil {
		log.Fatal(err)
	}
	table.SetRowLine(true)
	table.SetAutoMergeCellsByColumnIndex([]int{1, 2})
	table.Render()

	fmt.Println()
	fmt.Println()

	table.SetAutoMergeCellsByColumnIndex([]int{3})
	table.Render()
}
