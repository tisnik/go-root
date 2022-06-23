// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/

package main

import (
	"log"

	"go/ast"
	"go/parser"
	"go/token"
)

// zdrojový kód, který se má naparsovat
const source = `
package main

var answer int = 42
`

func main() {
	// struktura reprezentující množinu zdrojových kódů
	fileSet := token.NewFileSet()

	// konstrukce parseru a parsing zdrojového kódu
	f, err := parser.ParseFile(fileSet, "", source, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	// čitelný výpis obsahu abstraktího syntaktického stromu
	err = ast.Print(fileSet, f)
	if err != nil {
		panic(err)
	}
}
