// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá sedmá část
//    Skriptovací jazyk Lua v aplikacích naprogramovaných v Go
//    https://www.root.cz/clanky/skriptovaci-jazyk-lua-v-aplikacich-naprogramovanych-v-go/
//
// Seznam příkladů z dvacáté sedmé části:
//    https://github.com/tisnik/go-root/blob/master/article_27/README.md
//
// Demonstrační příklad číslo 2:
//    Poslání skriptu v řetězci do Lua VM.

package main

import (
	"github.com/yuin/gopher-lua"
	"log"
)

func main() {
	luaVM := lua.NewState()
	log.Println("Lua VM has been initialized")

	defer func() {
		luaVM.Close()
		log.Println("Lua VM has been closed")
	}()

	err := luaVM.DoString(`print("Hello from Lua!")`)
	if err != nil {
		log.Fatal(err)
	}
}
