package main

import (
	"fmt"

	"github.com/expr-lang/expr"
)

type User struct {
	Name string
}

func main() {
	user1 := new(User)
	user1.Name = "Pavel"

	user2 := new(User)
	user2.Name = "Petr"

	env := map[string]interface{}{
		"user1":   user1,
		"user2":   user2,
		"user3":   (*User)(nil),
		"println": fmt.Println,
	}

	code := `
	    println(user1?.Name);
	    println(user2?.Name);
	    println(user3?.Name);
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
