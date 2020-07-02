// Seriál "Programovací jazyk Go"
//
// Dvacátá sedmá část
//    Skriptovací jazyk Lua v aplikacích naprogramovaných v Go
//    https://www.root.cz/clanky/skriptovaci-jazyk-lua-v-aplikacich-naprogramovanych-v-go/
//
// Demonstrační příklad číslo 5:
//    Spuštění dvou Lua VM, každé v jedné gorutině.

package main

import (
	"github.com/yuin/gopher-lua"
	"log"
)

// LuaSource1 contains file name with first script written in Lua
const LuaSource1 = "l1.lua"

// LuaSource2 contains file name with second script written in Lua
const LuaSource2 = "l2.lua"

func callLuaVM1(c chan bool) {
	defer func() { c <- true }()

	luaVM1 := lua.NewState()
	log.Println("Lua VM1 has been initialized")

	defer func() {
		luaVM1.Close()
		log.Println("Lua VM1 has been closed")
	}()

	err := luaVM1.DoFile(LuaSource1)
	if err != nil {
		log.Fatal(err)
	}
}

func callLuaVM2(c chan bool) {
	defer func() { c <- true }()

	luaVM2 := lua.NewState()
	log.Println("Lua VM2 has been initialized")

	defer func() {
		luaVM2.Close()
		log.Println("Lua VM2 has been closed")
	}()

	err := luaVM2.DoFile(LuaSource2)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	c1 := make(chan bool)
	c2 := make(chan bool)

	go callLuaVM1(c1)
	go callLuaVM2(c2)

	<-c1
	<-c2
}
