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
	for i := uint8(0); i < 24; i++ {
		message := fmt.Sprintf("Grayscale %d", i)
		fmt.Println(colorizer.Gray(i, message))
	}
}
