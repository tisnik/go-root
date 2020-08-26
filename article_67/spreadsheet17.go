package main

import (
	"fmt"

	"github.com/tealeg/xlsx/v3"
)

// jméno generovaného souboru
const spreadsheetName = "test17.xlsx"

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
	cell.SetString("1x2")

	// naplnění buňky hodnotou
	cell.SetString("1x2")

	// a spojení s buňkou na následujícím řádku
	cell.Merge(0, 1)

	// přidání druhé buňky na řádek
	cell = row.AddCell()

	// naplnění buňky hodnotou
	cell.SetString("1x1")

	// další řádek tabulky
	sheet.AddRow().AddCell().SetString("foobar")

	// pokus o uložení sešitu
	err = worksheet.Save(spreadsheetName)
	if err != nil {
		panic(err)
	}
}
