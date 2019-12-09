// Seriál "Programovací jazyk Go"
//
// Dvacátá sedmá část
//
// Demonstrační příklad číslo 8:
//    Zavolání Lua funkce z Go s předáním parametrů.
//

package main

import (
	"github.com/yuin/gopher-lua"
	"log"
)

const LuaSource = "function_params.lua"

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
	}, lua.LString("foo"), lua.LString("bar"))

	if err != nil {
		log.Fatal(err)
	}

}
