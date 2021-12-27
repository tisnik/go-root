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

func main() {
    var x = 10
    var y = 20

    if x > 0 {
    }

    if x != y {
    }

    if 0 < x {
    }
}
`

func getValue(n ast.Expr) string {
	switch v := n.(type) {
	case *ast.BasicLit:
		return v.Value
	case *ast.Ident:
		return v.Name
	case *ast.ArrayType:
		if v.Len == nil {
			return fmt.Sprintf("Slice of %s\n", v.Elt)
		} else {
			return fmt.Sprintf("Array of %s\n", v.Elt)
		}
	default:
		return fmt.Sprintf("Unrecognized type %T\n", v)
	}
}

// funkce volaná při průchodu AST
func inspectCallback(n ast.Node) bool {
	// pokud se jedná o volání funkce, vrátí se hodnota + true
	ifStatement, found := n.(*ast.IfStmt)
	if found {
		fmt.Print("if statement")
		condition := ifStatement.Cond
		binaryExpr, found := condition.(*ast.BinaryExpr)
		if found {
			fmt.Print(" with binary condition")
			left := getValue(binaryExpr.X)
			right := getValue(binaryExpr.Y)
			operand := binaryExpr.Op
			fmt.Printf(" %s %s %s", left, operand, right)
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
