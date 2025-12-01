// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
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
var x2 []float = make([]float, 10)
var x3 []T = make([]T, 20)

var y = len(x)
var z = cap(x)
var w = len(x) + cap(x)
var a = append(x, 10)
var b = foo([10]int)
var c = bar(x, y, z, 1, 1.0, 1i, "foo", 'b')

var d = c.baz(1, 2, x)
`

// funkce volaná při průchodu AST
func inspectCallback(n ast.Node) bool {
	// pokud se jedná o volání funkce, vrátí se hodnota + true
	funcCall, ok := n.(*ast.CallExpr)
	if ok {
		method, ok := funcCall.Fun.(*ast.SelectorExpr)
		if ok {
			fmt.Printf("Method: %s for type %s ", method.Sel.Name, method.X)
		} else {
			fmt.Printf("Function: %s ", funcCall.Fun)
		}
		// výpis podrobnějších informací o volané funkci
		fmt.Printf("with %d arguments:\n", len(funcCall.Args))
		// výpis informací o argumentech funkce
		for i, arg := range funcCall.Args {
			fmt.Printf("\t%d\t", i+1)
			switch v := arg.(type) {
			case *ast.BasicLit:
				fmt.Printf("Constant %s of type %s\n", v.Value, v.Kind)
			case *ast.Ident:
				fmt.Printf("Variable %s\n", v.Name)
			case *ast.ArrayType:
				if v.Len == nil {
					fmt.Printf("Slice of %s\n", v.Elt)
				} else {
					fmt.Printf("Array of %s\n", v.Elt)
				}
			default:
				fmt.Printf("Unrecognized type %T\n", v)
			}
		}
		fmt.Println()
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
