// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá šestá část:
//    Jazyk Go a textový terminál ve funkci základního prvku uživatelského rozhraní
//    https://www.root.cz/clanky/jazyk-go-a-textovy-terminal-ve-funkci-zakladniho-prvku-uzivatelskeho-rozhrani/
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
	fmt.Println(colorizer.BgRed("Test"))
	fmt.Println(colorizer.BgGreen("Test"))
	fmt.Println(colorizer.BgBlue("Test"))
	fmt.Println(colorizer.BgCyan("Test"))
	fmt.Println(colorizer.BgMagenta("Test"))
	fmt.Println(colorizer.BgYellow("Test"))

	fmt.Println()

	fmt.Println(colorizer.BgRed("Test").Bold())
	fmt.Println(colorizer.BgGreen("Test").Bold())
	fmt.Println(colorizer.BgBlue("Test").Bold())
	fmt.Println(colorizer.BgCyan("Test").Bold())
	fmt.Println(colorizer.BgMagenta("Test").Bold())
	fmt.Println(colorizer.BgYellow("Test").Bold())
}
