package main

import (
	"fmt"

	"github.com/expr-lang/expr"
)

func main() {
	env := map[string]interface{}{
		"foo":     1,
		"bar":     false,
		"baz":     nil,
		"println": fmt.Println,
	}

	code := `
	    println("foo=", foo ?? "*** nil ***");
	    println("bar=", bar ?? "*** nil ***");
	    println("baz=", baz ?? "*** nil ***");
	`

	fmt.Println("Compiling:", code)
	program, err := expr.Compile(code, expr.Env(env))
	if err != nil {
		panic(err)
	}
	fmt.Println("Compiled")

	fmt.Println("Calling compiled program:")

	output, err := expr.Run(program, env)
	if err != nil {
		panic(err)
	}

	fmt.Println("Output:", output)
}
