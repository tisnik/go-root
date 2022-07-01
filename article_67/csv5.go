// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šedesátá sedmá část
//    Tvorba sešitů pro tabulkové procesory v programovacím jazyku Go
//    https://www.root.cz/clanky/tvorba-sesitu-pro-tabulkove-procesory-v-programovacim-jazyku-go/
//
// Seznam příkladů z šedesáté sedmé části:
//    https://github.com/tisnik/go-root/blob/master/article_67/README.md
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_67/csv5.html

package main

import (
	"encoding/csv"
	"os"
)

// jméno generovaného souboru
const spreadsheetName = "test5.csv"

func main() {
	records := [][]string{
		{"first\nname", "last\nname", "username"},
		{"'Rob'", "'Pike'", "rob"},
		{"\"Ken\"", "\"Thompson\"", "ken"},
		{"`Robert`", "`Griesemer`", "gri"},
		{"A B", "C\tD", "\n"},
		{"Foo,Bar", "Baz,,,", ","},
	}

	// vytvoření výstupního souboru s tabulkou
	fout, err := os.Create(spreadsheetName)
	if err != nil {
		panic(err)
	}

	// zajištění, že se soubor s tabulkou uzavře při ukončení programu
	defer fout.Close()

	// konstrukce objektu pro postupné vytvoření tabulky
	writer := csv.NewWriter(fout)

	for _, record := range records {
		// vložení nového řádku (záznamu)
		err := writer.Write(record)
		if err != nil {
			panic(err)
		}
	}

	// pokus o uložení tabulky a vyprázdnění bufferu
	writer.Flush()

	// test, zda při Write() nebo Flush() došlo k nějaké chybě
	err = writer.Error()
	if err != nil {
		panic(err)
	}
}
