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
//    https://tisnik.github.io/go-root/article_67/spreadsheet16.html

package main

import (
	"fmt"

	"github.com/tealeg/xlsx/v3"
)

// jméno generovaného souboru
const spreadsheetName = "test16.xlsx"

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

	// přidání první buňky na řádek
	cell := row.AddCell()

	// naplnění buňky hodnotou
	cell.SetString("1x1")

	// přidání druhé buňky na řádek
	cell = row.AddCell()

	// naplnění buňky hodnotou
	cell.SetString("2x1")

	// a spojení s další buňkou
	cell.Merge(1, 0)

	// pokus o uložení sešitu
	err = worksheet.Save(spreadsheetName)
	if err != nil {
		panic(err)
	}
}
