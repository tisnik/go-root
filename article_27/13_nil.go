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
