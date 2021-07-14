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
