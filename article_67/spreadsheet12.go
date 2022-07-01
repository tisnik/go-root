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
//    https://tisnik.github.io/go-root/article_67/spreadsheet12.html

package main

import (
	"fmt"
	"time"

	"github.com/tealeg/xlsx/v3"
)

// jméno generovaného souboru
const spreadsheetName = "test12.xlsx"

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

	// formáty časového razítka, které si odzkoušíme
	formats := []string{
		"yyyy-mm-dd",
		"m/d",
		"m/d/yyyy",
		"dd/mm/yyyy",
		"dd/mm/yy",
		"yyyy-mm-dd hh:mm:ss",
		"hh:mm",
		"hh:mm:ss",
		"h:mm:ss",
		"[h]:mm:ss",
		"[mm]:ss",
	}

	// aktuální datum a čas
	timestamp := time.Now()

	// zobrazit tabulku s daty/časy naformátovanými zvoleným formátem
	for _, format := range formats {
		// přidání nového řádku do tabulky
		row := sheet.AddRow()

		// popis formátu
		row.AddCell().SetString(format)

		// naformátované časové razítko
		cell := row.AddCell()
		cell.SetDateTime(timestamp)
		cell.SetFormat(format)
	}

	// pokus o uložení sešitu
	err = worksheet.Save(spreadsheetName)
	if err != nil {
		panic(err)
	}
}
