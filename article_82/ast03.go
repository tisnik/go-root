// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam příkladů z osmdesáté druhé části:
//    https://github.com/tisnik/go-root/blob/master/article_82/README.md

package main

import (
	"fmt"
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

// funkce volaná při průchodu AST
func inspectCallback(n ast.Node) bool {
	// dosáhli jsme koncového uzlu?
	if n == nil {
		return false
	}

	// tisk typu uzlu
	fmt.Printf("%T\n", n)
	return true
}

func main() {
	// struktura reprezentující množinu zdrojových kódů
	fileSet := token.NewFileSet()

	// konstrukce parseru a parsing zdrojového kódu
	f, err := parser.ParseFile(fileSet, "", source, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	// zahájení průchodu abstraktním syntaktickým stromem
	ast.Inspect(f, inspectCallback)
}
