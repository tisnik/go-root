package main

import (
	"fmt"
	"log"
	"strings"

	"go/ast"
	"go/parser"
)

// výraz, který se má naparsovat
const source = `
((1 + x * 2) * -(2 + y / -z)) % (x + -y + -z)
`

// nový datový typ implementující rozhraní ast.Visitor
type visitor int

// implementace (jediné) funkce předepsané v rozhraní ast.Visitor
func (v visitor) Visit(n ast.Node) ast.Visitor {
	// dosáhli jsme koncového uzlu?
	if n == nil {
		return nil
	}

	// tisk pozice a typu uzlu
	fmt.Printf("%3d\t", v)
	var s string

	switch x := n.(type) {
	case *ast.BasicLit:
		s = "Literal: " + x.Value
	case *ast.Ident:
		s = "Identifier: " + x.Name
	case *ast.UnaryExpr:
		s = "Unary operator: " + x.Op.String()
	case *ast.BinaryExpr:
		s = "Binary operator: " + x.Op.String()
	case *ast.ParenExpr:
		s = "Expression in parenthesis"
	}

	indent := strings.Repeat("\t", int(v))
	if s != "" {
		fmt.Printf("%s%s\n", indent, s)
	} else {
		fmt.Printf("%s%T\n", indent, n)
	}
	return v + 1
}

func main() {
	// konstrukce parseru a parsing zdrojového kódu
	f, err := parser.ParseExpr(source)
	if err != nil {
		log.Fatal(err)
	}

	var v visitor

	// zahájení průchodu abstraktním syntaktickým stromem
	ast.Walk(v, f)
}
