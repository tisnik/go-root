// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá sedmá část
//    Skriptovací jazyk Lua v aplikacích naprogramovaných v Go
//    https://www.root.cz/clanky/skriptovaci-jazyk-lua-v-aplikacich-naprogramovanych-v-go/
//
// Demonstrační příklad číslo 17:
//    Zjištění návratové hodnoty a jejího typu.

package main

import (
	"github.com/yuin/gopher-lua"
	"log"
)

const LuaSource = "go_from_lua_add.lua"

func compute(L *lua.LState) int {
	log.Println("called from Lua")
	a := L.ToInt(1)
	b := L.ToInt(2)
	log.Printf("1st parameter %d\n", a)
	log.Printf("2nd parameter %d\n", b)
	L.Push(lua.LNumber(a + b))
	return 1
}

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

	luaVM.SetGlobal("compute", luaVM.NewFunction(compute))

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
	println("Value", ret.String())
}
