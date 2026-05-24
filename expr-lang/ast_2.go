package main

import (
	"fmt"
	"log"

	"github.com/expr-lang/expr/ast"
	"github.com/expr-lang/expr/parser"
)

func main() {

	code := `1+2+3+3*(5+7)`

	tree, err := parser.Parse(code)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ast.Dump(tree.Node))
}
