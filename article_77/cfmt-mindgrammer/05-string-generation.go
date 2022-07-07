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

	msg = cfmt.Ssuccess("Success message")
	fmt.Printf("1st line: %s\n", msg)

	msg = cfmt.Sinfo("Info message")
	fmt.Printf("2nd line: %s\n", msg)

	msg = cfmt.Swarning("Warning message")
	fmt.Printf("3rd line: %s\n", msg)

	msg = cfmt.Serror("Error message")
	fmt.Printf("4th line: %s\n", msg)

	fmt.Println()
	fmt.Println("That's all...")
}
