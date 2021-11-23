// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šedesátá sedmá část
//    Tvorba sešitů pro tabulkové procesory v programovacím jazyku Go – formát xlsx
//    https://www.root.cz/clanky/tvorba-sesitu-pro-tabulkove-procesory-v-programovacim-jazyku-go-format-xlsx/
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_67/spreadsheet03.html

package main

import (
	"fmt"

	"github.com/tealeg/xlsx/v3"
)

// jméno generovaného souboru
const spreadsheetName = "test03.xlsx"

func main() {
	// konstrukce struktury typu File
	worksheet := xlsx.NewFile()

	fmt.Println(worksheet)

	// přidání tří listů do sešitu
	names := []string{"Tabulka1", "Tabulka2", "Tabulka3"}
	for _, name := range names {
		// pokus o přidání nového listu
		sheet, err := worksheet.AddSheet(name)
		if err != nil {
			panic(err)
		}

		fmt.Println(sheet)
	}

	// pokus o uložení sešitu
	err := worksheet.Save(spreadsheetName)
	if err != nil {
		panic(err)
	}
}
