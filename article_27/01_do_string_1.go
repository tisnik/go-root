// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá sedmá část
//    Skriptovací jazyk Lua v aplikacích naprogramovaných v Go
//    https://www.root.cz/clanky/skriptovaci-jazyk-lua-v-aplikacich-naprogramovanych-v-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z dvacáté sedmé části:
//    https://github.com/tisnik/go-root/blob/master/article_27/README.md
//
// Demonstrační příklad číslo 1:
//    Vytvoření virtuálního stroje jazyka Lua, poslání skriptu v řetězci.

package main

import (
	"github.com/yuin/gopher-lua"
	"log"
)

func main() {
	luaVM := lua.NewState()

	defer luaVM.Close()

	err := luaVM.DoString(`print("Hello from Lua")`)
	if err != nil {
		log.Fatal(err)
	}
}
