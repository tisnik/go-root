package main

import (
	"fmt"
	"log"

	"github.com/expr-lang/expr/ast"
	"github.com/expr-lang/expr/parser"
)

func main() {
	code := `
            if foo < 0 {
                let result = -foo*bar;
	        result
	    } else {
	        let result = foo*bar;
	        result
	    }`

	tree, err := parser.Parse(code)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ast.Dump(tree.Node))
}
