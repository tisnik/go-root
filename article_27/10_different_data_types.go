// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá sedmá část
//    Skriptovací jazyk Lua v aplikacích naprogramovaných v Go
//    https://www.root.cz/clanky/skriptovaci-jazyk-lua-v-aplikacich-naprogramovanych-v-go/
//
// Demonstrační příklad číslo 10:
//    Použití datových typů jazyka Lua.

package main

import (
	"github.com/yuin/gopher-lua"
	"log"
)

// LuaSource contains file name with script written in Lua
const LuaSource = "concatenate.lua"

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
		Fn:      luaVM.GetGlobal("concatenate"),
		NRet:    1,
		Protect: true,
	}, lua.LString("foo"), lua.LString("bar"))

	if err != nil {
		log.Fatal(err)
	}

	ret := luaVM.Get(-1)
	luaVM.Pop(1)
	println("Type", ret.Type())
	println("Value", ret.String())
}
