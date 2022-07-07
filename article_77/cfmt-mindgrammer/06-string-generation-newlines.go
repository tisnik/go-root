// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá sedmá část:
//    Jazyk Go a textový terminál ve funkci základního prvku uživatelského rozhraní (2.část)
//    https://www.root.cz/clanky/jazyk-go-a-textovy-terminal-ve-funkci-zakladniho-prvku-uzivatelskeho-rozhrani-2-cast/
//
// Seznam příkladů ze sedmdesáté sedmé části:
//    https://github.com/tisnik/go-root/blob/master/article_77/README.md

package main

import (
	"fmt"

	"github.com/mingrammer/cfmt"
)

func main() {
	var msg string

	msg = cfmt.Ssuccessln("Success message")
	fmt.Printf("1st line: %s", msg)

	msg = cfmt.Sinfoln("Info message")
	fmt.Printf("2nd line: %s", msg)

	msg = cfmt.Swarningln("Warning message")
	fmt.Printf("3rd line: %s", msg)

	msg = cfmt.Serrorln("Error message")
	fmt.Printf("4th line: %s", msg)

	fmt.Println()
	fmt.Println("That's all...")
}
