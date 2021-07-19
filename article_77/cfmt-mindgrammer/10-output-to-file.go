package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mingrammer/cfmt"
)

func main() {
	file, err := os.Create("./temp.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	for i := 1; i <= 10; i++ {
		switch {
		case i < 5:
			cfmt.Fwarningf(file, "Value too low: %d\n", i)
		case i == 5:
			cfmt.Fsuccessf(file, "An ideal value: %d\n", i)
		case i == 10:
			cfmt.Ferrorf(file, "Too high!!! %d\n", i)
		case i > 5:
			cfmt.Finfof(file, "Bit higher then necessary: %d\n", i)
		default:
			cfmt.Ferrorf(file, "Can't happen %d\n", i)
		}
	}

	fmt.Println()
	fmt.Println("That's all...")
}
