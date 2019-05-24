package main

import (
	"github.com/yuin/gopher-lua"
	"log"
)

func main() {
	luaVM := lua.NewState()

	defer luaVM.Close()

	err := luaVM.DoString(`print("Hello from Lua")`)
	if err != nil {
		log.Fatal(err)
	}
}
