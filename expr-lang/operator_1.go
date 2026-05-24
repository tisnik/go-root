package main

import (
	"fmt"

	"github.com/expr-lang/expr"
)

func mul(x, y int) int {
	return 2 * x * y
}

func main() {
	env := map[string]interface{}{
		"foo": 6,
		"bar": 7,
		"mul": mul,
	}

	code := `foo*bar`

	fmt.Println("Compiling:", code)
	program, err := expr.Compile(code, expr.Env(env), expr.Operator("*", "mul"))
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
