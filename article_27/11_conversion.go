// Seriál "Programovací jazyk Go"
//
// Dvacátá sedmá část
//
// Demonstrační příklad číslo 11:
//    Konverze datových typů jazyka Lua.

package main

import (
	"github.com/yuin/gopher-lua"
	"log"
)

// LuaSource contains file name with script written in Lua
const LuaSource = "add.lua"

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
		Fn:      luaVM.GetGlobal("add"),
		NRet:    1,
		Protect: true,
	}, lua.LNumber(1), lua.LNumber(2))

	if err != nil {
		log.Fatal(err)
	}

	ret := luaVM.Get(-1)
	luaVM.Pop(1)
	println("Type", ret.Type())
	if number, ok := ret.(lua.LNumber); ok {
		println("Value", float64(number))
	}
}
