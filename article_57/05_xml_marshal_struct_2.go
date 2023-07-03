// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá sedmá část
//    Serializace a deserializace datových struktur v programovacím jazyce Go
//    https://www.root.cz/clanky/serializace-a-deserializace-datovych-struktur-v-programovacim-jazyce-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z padesáté sedmé části:
//    https://github.com/tisnik/go-root/blob/master/article_57/README.md

package main

import (
	"encoding/xml"
	"fmt"
)

type User1 struct {
	id      uint32 `xml:"id"`
	name    string `xml:"user_name"`
	surname string `xml:"surname"`
}

type User2 struct {
	Id      uint32 `xml:"id"`
	Name    string `xml:"user_name"`
	Surname string `xml:"surname"`
}

func main() {
	user1 := User1{
		1,
		"Pepek",
		"Vyskoč"}

	user2 := User2{
		1,
		"Pepek",
		"Vyskoč"}

	user1asXML, _ := xml.Marshal(user1)
	fmt.Println(string(user1asXML))

	fmt.Println()

	user2asXML, _ := xml.Marshal(user2)
	fmt.Println(string(user2asXML))
}
