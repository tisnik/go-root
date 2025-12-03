// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam příkladů z osmdesáté třetí části:
//    https://github.com/tisnik/go-root/blob/master/article_83/README.md
//

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// zdrojový kód, který se má naparsovat
const source = `
package main

var answer int = 42
var x []int = make([]int, 10)
var y = len(x)
var z = cap(x)
var w = len(x) + cap(x)
var a = append(x, 10)
`

// funkce volaná při průchodu AST
func inspectCallback(n ast.Node) bool {
	// pokud se jedná o volání funkce, vrátí se hodnota + true
	funcCall, ok := n.(*ast.CallExpr)
	if ok {
		// výpis základní informace o volané funkci
		fmt.Println(funcCall.Fun)
	}
	return true
}

func main() {
	// struktura reprezentující množinu zdrojových kódů
	fileSet := token.NewFileSet()

	// konstrukce parseru a parsing zdrojového kódu
	file, err := parser.ParseFile(fileSet, "", source, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// zahájení průchodu abstraktním syntaktickým stromem
	ast.Inspect(file, inspectCallback)
}
