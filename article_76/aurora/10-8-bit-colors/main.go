// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá šestá část:
//    Jazyk Go a textový terminál ve funkci základního prvku uživatelského rozhraní
//    https://www.root.cz/clanky/jazyk-go-a-textovy-terminal-ve-funkci-zakladniho-prvku-uzivatelskeho-rozhrani/

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
	for j := uint8(16); j <= 231; j += 4 {
		for i := uint8(16); i < 231; i += 2 {
			fmt.Print(colorizer.Index(i, colorizer.BgIndex(j, "x")))
		}
		fmt.Println()
	}
}
