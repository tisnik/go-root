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
	for j := uint8(24); j > 0; j-- {
		for i := uint8(0); i < 24; i += 2 {
			message := fmt.Sprintf(" [%2d %2d] ", i, j)
			fmt.Print(colorizer.BgGray(j, colorizer.Gray(i, message)))
		}
		fmt.Println()
	}
}
