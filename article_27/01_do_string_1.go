// Seriál "Programovací jazyk Go"
//
// Dvacátá sedmá část
//    Skriptovací jazyk Lua v aplikacích naprogramovaných v Go
//    https://www.root.cz/clanky/skriptovaci-jazyk-lua-v-aplikacich-naprogramovanych-v-go/
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
