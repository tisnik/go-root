package main

import (
	"fmt"

	"github.com/expr-lang/expr"
)

func main() {
	env := map[string]interface{}{
		"foo": -6,
		"bar": 7,
	}

	code := `
            if foo < 0 {
                let result = -foo*bar;
	        result
	    } else {
	        let result = foo*bar;
	        result
	    }`

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
