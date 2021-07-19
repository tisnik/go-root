package main

import (
	"fmt"

	"github.com/mingrammer/cfmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		switch {
		case i < 5:
			cfmt.Warningf("Value too low: %d\n", i)
		case i == 5:
			cfmt.Successf("An ideal value: %d\n", i)
		case i == 10:
			cfmt.Errorf("Too high!!! %d\n", i)
		case i > 5:
			cfmt.Infof("Bit higher then necessary: %d\n", i)
		default:
			cfmt.Errorf("Can't happen %d\n", i)
		}
	}

	fmt.Println()
	fmt.Println("That's all...")
}
