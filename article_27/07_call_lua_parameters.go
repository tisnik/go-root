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
	}, lua.LNumber(1), lua.LNumber(2))

	if err != nil {
		log.Fatal(err)
	}

}
