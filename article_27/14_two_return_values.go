// Seriál "Programovací jazyk Go"
//
// Dvacátá sedmá část
//
// Demonstrační příklad číslo 14:
//    Funkce vracející dvě hodnoty.

package main

import (
	"github.com/yuin/gopher-lua"
	"log"
)

const LuaSource = "swap.lua"

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
		Fn:      luaVM.GetGlobal("swap"),
		NRet:    2,
		Protect: true,
	}, lua.LNumber(1), lua.LNumber(2))

	if err != nil {
		log.Fatal(err)
	}

	ret1 := luaVM.Get(-2)
	ret2 := luaVM.Get(-1)
	luaVM.Pop(2)
	println("Type", ret1.Type())
	println("Type", ret2.Type())
	if number, ok := ret1.(lua.LNumber); ok {
		println("Value #1", float64(number))
	}
	if number, ok := ret2.(lua.LNumber); ok {
		println("Value #2", float64(number))
	}
}
