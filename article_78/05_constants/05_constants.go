// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Sedmdesátá osmá část:
//    Jazyk Go a vyhodnocování výrazů v době běhu aplikace
//    https://www.root.cz/clanky/jazyk-go-a-vyhodnocovani-vyrazu-v-dobe-behu-aplikace/
//
// Seznam příkladů ze sedmdesáté osmé části:
//    https://github.com/tisnik/go-root/blob/master/article_78/README.md

package main

import (
	"fmt"
	"os"

	"github.com/PaesslerAG/gval"
)

func main() {
	// parametry předávané vyhodnocovanému výrazu
	parameters := map[string]interface{}{
		"x": 10,
		"y": 20,
	}

	// vyhodnocení výrazu
	value, err := gval.Evaluate("2*x + y", parameters)
	if err != nil {
		// kód pro zpracování chyby při vyhodnocování výrazu
		fmt.Println(err)
		os.Exit(1)
	}

	// výpis výsledku výrazu
	fmt.Print(value)
}
