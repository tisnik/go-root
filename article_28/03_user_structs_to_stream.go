// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá osmá část
//    Datové struktury s líným vyhodnocováním v programovacím jazyce Go
//    https://www.root.cz/clanky/datove-struktury-s-linym-vyhodnocovanim-v-programovacim-jazyce-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z dvacáté osmé části:
//    https://github.com/tisnik/go-root/blob/master/article_28/README.md
//
// Demonstrační příklad číslo 3:
//    Struktury konvertované do streamů.

package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

// User data type represents an user in (some) information system
type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	var users = []User{
		{
			id:      1,
			name:    "Pepek",
			surname: "Vyskoč"},
		{
			id:      2,
			name:    "Pepek",
			surname: "Vyskoč"},
		{
			id:      3,
			name:    "Josef",
			surname: "Vyskočil"},
	}
	fmt.Println(users)

	fmt.Printf("input:  %v\n", users)

	stream := koazee.StreamOf(users)
	fmt.Printf("stream: %v\n", stream.Out().Val())
}
