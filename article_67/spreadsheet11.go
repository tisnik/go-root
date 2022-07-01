// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šedesátá sedmá část
//    Tvorba sešitů pro tabulkové procesory v programovacím jazyku Go – formát xlsx
//    https://www.root.cz/clanky/tvorba-sesitu-pro-tabulkove-procesory-v-programovacim-jazyku-go-format-xlsx/
//
// Seznam příkladů z šedesáté sedmé části:
//    https://github.com/tisnik/go-root/blob/master/article_67/README.md
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_67/spreadsheet11.html

package main

import (
	"fmt"

	"github.com/tealeg/xlsx/v3"
)

// jméno generovaného souboru
const spreadsheetName = "test11.xlsx"

func main() {
	// konstrukce struktury typu File
	worksheet := xlsx.NewFile()

	fmt.Println(worksheet)

	// přidání listu do sešitu
	sheet, err := worksheet.AddSheet("Tabulka1")
	if err != nil {
		panic(err)
	}
	defer sheet.Close()

	fmt.Println(sheet)

	// přidání řádku do tabulky
	row := sheet.AddRow()

	style := xlsx.NewStyle()
	style.Alignment.Horizontal = "right"
	style.Fill.FgColor = "ffa0ffa0"
	style.Fill.PatternType = "solid"
	style.Font.Bold = true
	style.ApplyAlignment = true
	style.ApplyFill = true
	style.ApplyFont = true

	// přidání první buňky na řádek a naplnění hodnotou
	cell := row.AddCell()
	cell.SetString("Test #1")

	// nastavení stylu buňky
	cell.SetStyle(style)

	style.Fill.FgColor = "ffffa0a0"
	style.Fill.PatternType = "solid"
	style.Font.Bold = false

	// přidání druhé buňky na řádek a naplnění hodnotou
	cell = row.AddCell()
	cell.SetString("Test #1")

	// nastavení stylu buňky
	cell.SetStyle(style)

	// pokus o uložení sešitu
	err = worksheet.Save(spreadsheetName)
	if err != nil {
		panic(err)
	}
}
