// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze třetí části:
//    https://github.com/tisnik/go-root/blob/master/article_03/README.md
//
// Demonstrační příklad číslo 21:
//    Odstranění položek z mapy.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_03/21_delete_from_map.html

package main

import "fmt"

// Key je uživatelsky definovaná datová struktura
type Key struct {
	id   uint32
	role string
}

// User je uživatelsky definovaná datová struktura
type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	m1 := make(map[Key]User)
	fmt.Println(m1)

	m1[Key{1, "admin"}] = User{
		id:      1,
		name:    "Pepek",
		surname: "Vyskoč"}

	m1[Key{2, "user"}] = User{
		id:      2,
		name:    "Josef",
		surname: "Vyskočil"}

	fmt.Println(m1)

	delete(m1, Key{2, "user"})
	fmt.Println(m1)

	delete(m1, Key{1000, "nekdo jiny"})
	fmt.Println(m1)
}
