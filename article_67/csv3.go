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
//    https://tisnik.github.io/go-root/article_67/csv3.html

package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

// jméno generovaného souboru
const spreadsheetName = "test3.csv"

func factorial(n int64) int64 {
	switch {
	case n < 0:
		return 1
	case n == 0:
		return 1
	default:
		return n * factorial(n-1)
	}
}

func main() {
	// vytvoření výstupního souboru s tabulkou
	fout, err := os.Create(spreadsheetName)
	if err != nil {
		panic(err)
	}

	// zajištění, že se soubor s tabulkou uzavře při ukončení programu
	defer fout.Close()

	// konstrukce objektu pro postupné vytvoření tabulky
	writer := csv.NewWriter(fout)

	// zápis hlavičky
	header := []string{"n", "n!"}
	err = writer.Write(header)
	if err != nil {
		panic(err)
	}

	var records [][]string

	// zápis ostatních řádků
	for n := int64(0); n <= 20; n++ {
		// výpočet faktoriálu
		f := factorial(n)

		// příprava dat pro zápis
		var record []string
		record = append(record, strconv.FormatInt(n, 10), strconv.FormatInt(f, 10))

		// připojení k výsledkům
		records = append(records, record)
	}

	// uložení všech řádků (záznamu) a provedení operace Flush
	writer.WriteAll(records)

	// test, zda při Write() nebo Flush() došlo k nějaké chybě
	err = writer.Error()
	if err != nil {
		panic(err)
	}
}
