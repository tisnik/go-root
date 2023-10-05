// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá šestá část:
//    Jazyk Go a textový terminál ve funkci základního prvku uživatelského rozhraní
//    https://www.root.cz/clanky/jazyk-go-a-textovy-terminal-ve-funkci-zakladniho-prvku-uzivatelskeho-rozhrani/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam příkladů ze sedmdesáté šesté části:
//    https://github.com/tisnik/go-root/blob/master/article_76/README.md

package main

import (
	"flag"
	"fmt"

	"github.com/logrusorgru/aurora"
)

var colorizer aurora.Aurora

func init() {
	var colors = flag.Bool("colors", false, "enable or disable colors")
	flag.Parse()

	colorizer = aurora.NewAurora(*colors)
}

func main() {
	fmt.Println(colorizer.Red("Test"))
	fmt.Println(colorizer.Green("Test"))
	fmt.Println(colorizer.Blue("Test"))
	fmt.Println(colorizer.Cyan("Test"))
	fmt.Println(colorizer.Magenta("Test"))
	fmt.Println(colorizer.Yellow("Test"))

	fmt.Println()

	fmt.Println(colorizer.Red("Test").Bold())
	fmt.Println(colorizer.Green("Test").Bold())
	fmt.Println(colorizer.Blue("Test").Bold())
	fmt.Println(colorizer.Cyan("Test").Bold())
	fmt.Println(colorizer.Magenta("Test").Bold())
	fmt.Println(colorizer.Yellow("Test").Bold())
}
