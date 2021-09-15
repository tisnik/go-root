package main

import (
	"fmt"
	"os"

	"github.com/PaesslerAG/gval"
)

func main() {
	// parametry předávané vyhodnocovanému výrazu
	parameters := make(map[string]interface{})

	// vyhodnocení výrazu
	value, err := gval.Evaluate("6*7", parameters)
	if err != nil {
		// kód pro zpracování chyby při vyhodnocování výrazu
		fmt.Println(err)
		os.Exit(1)
	}

	// výpis výsledku výrazu
	fmt.Print(value)
}
