// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá sedmá část:
//    Jazyk Go a textový terminál ve funkci základního prvku uživatelského rozhraní (2.část)
//    https://www.root.cz/clanky/jazyk-go-a-textovy-terminal-ve-funkci-zakladniho-prvku-uzivatelskeho-rozhrani-2-cast/

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
