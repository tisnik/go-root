package main

import (
	"github.com/yuin/gopher-lua"
	"log"
)

const LuaSource = "go_from_lua.lua"

func Hello(L *lua.LState) int {
	log.Println("Hello from Go!")
	return 0
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

	luaVM.SetGlobal("hello", luaVM.NewFunction(Hello))

	err = luaVM.CallByParam(lua.P{
		Fn:      luaVM.GetGlobal("call_go"),
		NRet:    0,
		Protect: true,
	})

	if err != nil {
		log.Fatal(err)
	}
}
