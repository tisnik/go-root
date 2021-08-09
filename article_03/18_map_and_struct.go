// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Demonstrační příklad číslo 18:
//    Mapa struktur.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_03/18_map_and_struct.html

package main

import "fmt"

// User je uživatelsky definovaná datová struktura
type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	m1 := make(map[string]User)
	fmt.Println(m1)

	m1["prvni"] = User{
		id:      1,
		name:    "Pepek",
		surname: "Vyskoč"}

	m1["druhy"] = User{
		id:      2,
		name:    "Josef",
		surname: "Vyskočil"}

	fmt.Println(m1)
}
