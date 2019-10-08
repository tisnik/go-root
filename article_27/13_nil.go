// Seriál "Programovací jazyk Go"
//
// Dvacátá sedmá část
//
// Demonstrační příklad číslo 13:
//    Hodnota nil.

package main

import (
	"github.com/yuin/gopher-lua"
	"log"
)

const LuaSource = "return_nil.lua"

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
		Fn:      luaVM.GetGlobal("return_nil"),
		NRet:    1,
		Protect: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	ret := luaVM.Get(-1)
	luaVM.Pop(1)
	println("Type", ret.Type())
	println("Value", ret.String())
	println("is nil?", ret == lua.LNil)
}
