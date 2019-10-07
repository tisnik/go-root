// Seriál "Programovací jazyk Go"
//
// Dvacátá sedmá část
//
// Demonstrační příklad číslo 6:
//    Zavolání Lua funkce z Go.

package main

import (
	"github.com/yuin/gopher-lua"
	"log"
)

const LuaSource = "function.lua"

func main() {
	luaVM := lua.NewState()
	log.Println("Lua VM has been initialized")

	defer func() {
		luaVM.Close()
		log.Println("Lua VM has been closed")
	}()

	err := luaVM.DoFile(LuaSource)
	if err != nil {
		log.Fatal(err)
	}

	err = luaVM.CallByParam(lua.P{
		Fn:   luaVM.GetGlobal("hello"),
		NRet: 0,
	})

	if err != nil {
		log.Fatal(err)
	}

}
