package main

import (
	"fmt"

	"github.com/expr-lang/expr"
)

func main() {
	env := map[string]interface{}{
		"name": "Go",
		"version": map[string]int{
			"major": 1,
			"minor": 26,
			"patch": 3,
		},
		"println": fmt.Println,
	}

	code := `
            let major = string(version["major"]);
            let minor = string(version["minor"]);
            let patch = string(version["patch"]);
            let msg = name + " " + major + "." + minor + "." + patch;
	    println(msg)
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
