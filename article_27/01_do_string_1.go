// Seriál "Programovací jazyk Go"
//
// Dvacátá sedmá část
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
