package main

import (
	"fmt"

	"github.com/expr-lang/expr"
)

func compare(x, y int) bool {
	fmt.Printf("Comparing %d with %d\n", x, y)
	return x < y
}

func main() {
	env := map[string]interface{}{
		"foo": 6,
		"bar": 7,
		"lt":  compare,
	}

	code := `foo<bar`

	fmt.Println("Compiling:", code)
	program, err := expr.Compile(code, expr.Env(env), expr.Operator("<", "lt"))
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
