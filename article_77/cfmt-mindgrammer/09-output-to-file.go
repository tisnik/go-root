package main

import (
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

	cfmt.Fsuccessln(file, "Success message")
	cfmt.Finfoln(file, "Info message")
	cfmt.Fwarningln(file, "Warning message")
	cfmt.Ferrorln(file, "Error message")
}
