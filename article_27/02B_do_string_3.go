package main

import (
	"github.com/yuin/gopher-lua"
	"log"
)

const Script = `
for i = 1,10 do
    print("Hello from Lua!", i)
end`

func main() {
	luaVM := lua.NewState()
	log.Println("Lua VM has been initialized")

	defer func() {
		luaVM.Close()
		log.Println("Lua VM has been closed")
	}()

	log.Println("Starting the following Lua script:")
	log.Println(Script)

	err := luaVM.DoString(Script)
	if err != nil {
		log.Fatal(err)
	}
}
