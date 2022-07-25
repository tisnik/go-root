// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Seznam příkladů z osmdesáté druhé části:
//    https://github.com/tisnik/go-root/blob/master/article_82/README.md

package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

// zdrojový kód, který má být tokenizován
const source = `
var x const < 1 && + 2 * || 3 // nejlepší kód
`

func main() {
	// objekt představující scanner
	var s scanner.Scanner

	// struktura reprezentující množinu zdrojových kódů
	fset := token.NewFileSet()

	// přidání informace o zdrojovém kódu
	file := fset.AddFile("", fset.Base(), len(source))

	// inicializace scanneru
	s.Init(file, []byte(source), nil, scanner.ScanComments)

	// postupné provádění tokenizace a výpis jednotlivých tokenů
	for {
		pos, tok, lit := s.Scan()
		// byl nalezen speciální token reprezentující konec tokenizace
		if tok == token.EOF {
			fmt.Println("<EOF>")
			break
		}
		// výpis obsahu tokenu
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}

}
