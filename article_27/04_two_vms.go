package main

import (
	"github.com/yuin/gopher-lua"
	"log"
)

const LuaSource1 = "v1.lua"
const LuaSource2 = "v2.lua"
const LuaSource3 = "v3.lua"

func main() {
	luaVM1 := lua.NewState()
	log.Println("Lua VM1 has been initialized")

	luaVM2 := lua.NewState()
	log.Println("Lua VM2 has been initialized")

	defer func() {
		luaVM1.Close()
		log.Println("Lua VM1 has been closed")
	}()

	defer func() {
		luaVM2.Close()
		log.Println("Lua VM2 has been closed")
	}()

	err := luaVM1.DoFile(LuaSource1)
	if err != nil {
		log.Fatal(err)
	}

	err = luaVM2.DoFile(LuaSource2)
	if err != nil {
		log.Fatal(err)
	}

	err = luaVM1.DoFile(LuaSource3)
	if err != nil {
		log.Fatal(err)
	}
}
