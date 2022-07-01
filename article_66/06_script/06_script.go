// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šedesátá šestá část
//    Použití databáze Redis v aplikacích naprogramovaných v Go (2)
//    https://www.root.cz/clanky/pouziti-databaze-redis-v-aplikacich-naprogramovanych-v-go-2/
//
// Seznam příkladů z šedesáté šesté části:
//    https://github.com/tisnik/go-root/blob/master/article_66/README.md

package main

import (
	"fmt"
	"io/ioutil"

	"github.com/go-redis/redis/v8"
)

// adresa určující službu Redisu, která se má použít
const redisAddress = "localhost:6379"

const counterKey = "counter2"

func mustLoadScriptSource(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func main() {
	// vytvoření nového klienta s předáním konfiguračních parametrů
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// neměli bychom zapomenout na ukončení práce s klientem
	defer func() {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}()

	// získáme kontext
	context := client.Context()

	// pokus o klasický handshake typu PING-PONG
	_, err := client.Ping(context).Result()
	if err != nil {
		panic(err)
	}

	scriptSource := mustLoadScriptSource("script.lua")
	fmt.Printf("Loaded script:\n%s\n\n", scriptSource)

	script := redis.NewScript(scriptSource)

	n, err := script.Run(context, client, []string{counterKey}, 2).Result()
	fmt.Println(n, err)

	err = client.Set(context, counterKey, "40", 0).Err()
	if err != nil {
		panic(err)
	}

	n, err = script.Run(context, client, []string{counterKey}, 2).Result()
	fmt.Println(n, err)
}
