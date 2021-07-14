package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

func main() {
	var colorizer aurora.Aurora
	colorizer = aurora.NewAurora(true)

	fmt.Println(colorizer.Red("Test"))
	fmt.Println(colorizer.Green("Test"))
	fmt.Println(colorizer.Blue("Test"))
	fmt.Println(colorizer.Cyan("Test"))
	fmt.Println(colorizer.Magenta("Test"))
	fmt.Println(colorizer.Yellow("Test"))
}
