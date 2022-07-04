// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá první část
//    Formátovaný tisk obsahu tabulek a dalších datových struktur v Go
//    https://www.root.cz/clanky/formatovany-tisk-obsahu-tabulek-a-dalsich-datovych-struktur-v-go/

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
	table.SetBorder(false)

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgBlackColor},
		tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgWhiteColor})

	table.SetColumnColor(
		tablewriter.Colors{tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.FgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgWhiteColor})

	table.Render()
}
