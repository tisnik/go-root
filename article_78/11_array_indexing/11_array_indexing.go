package main

import (
	"fmt"
	"os"

	"github.com/PaesslerAG/gval"
)

func main() {
	// parametry předávané vyhodnocovanému výrazu
	parameters := map[string]interface{}{
		"x":   10,
		"y":   20,
		"arr": []int{10, 20, 30},
	}

	// vyhodnocení výrazu
	value, err := gval.Evaluate("arr[0] + arr[1] + arr[2]", parameters)
	if err != nil {
		// kód pro zpracování chyby při vyhodnocování výrazu
		fmt.Println(err)
		os.Exit(1)
	}

	// výpis výsledku výrazu
	fmt.Print(value)
}
