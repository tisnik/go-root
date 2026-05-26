package main

import (
	"fmt"
	"log"

	"github.com/expr-lang/expr/ast"
	"github.com/expr-lang/expr/parser"
)

func main() {
	code := `foo < 0 ? -foo*bar : foo*bar`

	tree, err := parser.Parse(code)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ast.Dump(tree.Node))
}
