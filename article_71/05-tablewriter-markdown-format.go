package main

import (
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	table, err := tablewriter.NewCSV(os.Stdout, "tiobe.csv", true)
	if err != nil {
		log.Fatal(err)
	}
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.Render()
}
