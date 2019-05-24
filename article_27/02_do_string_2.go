package main

import (
	"github.com/yuin/gopher-lua"
	"log"
)

func main() {
	luaVM := lua.NewState()
	log.Println("Lua VM has been initialized")

	defer func() {
		luaVM.Close()
		log.Println("Lua VM has been closed")
	}()

	err := luaVM.DoString(`print("Hello from Lua!")`)
	if err != nil {
		log.Fatal(err)
	}
}
