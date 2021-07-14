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

	fmt.Println(colorizer.Bold(colorizer.Red("Test")))
	fmt.Println(colorizer.Bold(colorizer.Green("Test")))
	fmt.Println(colorizer.Bold(colorizer.Blue("Test")))
	fmt.Println(colorizer.Bold(colorizer.Cyan("Test")))
	fmt.Println(colorizer.Bold(colorizer.Magenta("Test")))
	fmt.Println(colorizer.Bold(colorizer.Yellow("Test")))
}
