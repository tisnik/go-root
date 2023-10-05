// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá sedmá část:
//    Jazyk Go a textový terminál ve funkci základního prvku uživatelského rozhraní (2.část)
//    https://www.root.cz/clanky/jazyk-go-a-textovy-terminal-ve-funkci-zakladniho-prvku-uzivatelskeho-rozhrani-2-cast/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam příkladů ze sedmdesáté sedmé části:
//    https://github.com/tisnik/go-root/blob/master/article_77/README.md

package main

import (
	"encoding/hex"
	"fmt"

	"github.com/mingrammer/cfmt"
)

func printEncoded(s string) {
	bytes := []byte(s)
	fmt.Printf("Encoded:\n%s\n", hex.Dump(bytes))
}

func main() {
	var msg string

	msg = cfmt.Ssuccessln("Success message")
	fmt.Printf("1st line: %s", msg)
	printEncoded(msg)

	msg = cfmt.Sinfoln("Info message")
	fmt.Printf("2nd line: %s", msg)
	printEncoded(msg)

	msg = cfmt.Swarningln("Warning message")
	fmt.Printf("3rd line: %s", msg)
	printEncoded(msg)

	msg = cfmt.Serrorln("Error message")
	fmt.Printf("4th line: %s", msg)
	printEncoded(msg)

	fmt.Println()
	fmt.Println("That's all...")
}
