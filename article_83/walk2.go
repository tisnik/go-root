// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Seznam příkladů z osmdesáté třetí části:
//    https://github.com/tisnik/go-root/blob/master/article_83/README.md
//

package main

import (
	"fmt"
	"log"

	"go/ast"
	"go/parser"
)

// výraz, který se má naparsovat
const source = `
1 + 2 * 3 + x + y * z - 1
`

func walk(node ast.Node) {
	// dosáhli jsme koncového uzlu?
	if node == nil {
		return
	}

	var s string

	// převod uzlu do tisknutelné podoby a rekurzivní průchod poduzly
	switch x := node.(type) {
	case *ast.BasicLit:
		s = x.Value
	case *ast.Ident:
		s = x.Name
	case *ast.UnaryExpr:
		s = x.Op.String()
		walk(x.X)
	case *ast.BinaryExpr:
		walk(x.X)
		s = x.Op.String()
		walk(x.Y)
	}

	// tisk obsahu uzlu
	if s != "" {
		fmt.Printf("%s ", s)
	} else {
		fmt.Printf("%T ", node)
	}
}

func main() {
	// konstrukce parseru a parsing zdrojového kódu
	node, err := parser.ParseExpr(source)
	if err != nil {
		log.Fatal(err)
	}

	// zahájení průchodu abstraktním syntaktickým stromem
	walk(node)
}
