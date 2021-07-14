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
