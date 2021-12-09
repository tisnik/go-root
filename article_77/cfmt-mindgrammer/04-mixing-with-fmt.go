// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá sedmá část:
//    Jazyk Go a textový terminál ve funkci základního prvku uživatelského rozhraní (2.část)
//    https://www.root.cz/clanky/jazyk-go-a-textovy-terminal-ve-funkci-zakladniho-prvku-uzivatelskeho-rozhrani-2-cast/

package main

import (
	"fmt"

	"github.com/mingrammer/cfmt"
)

func main() {
	fmt.Print("1st line: ")
	cfmt.Successln("Success message")

	fmt.Print("2nd line: ")
	cfmt.Infoln("Info message")

	fmt.Print("3rd line: ")
	cfmt.Warningln("Warning message")

	fmt.Print("4th line: ")
	cfmt.Errorln("Error message")

	fmt.Println()
	fmt.Println("That's all...")
}
