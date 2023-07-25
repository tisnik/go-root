// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šedesátá pátá část
//    Použití databáze Redis v aplikacích naprogramovaných v Go
//    https://www.root.cz/clanky/pouziti-databaze-redis-v-aplikacich-naprogramovanych-v-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z šedesáté páté části:
//    https://github.com/tisnik/go-root/blob/master/article_65/README.md

package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// adresa určující službu Redisu, která se má použít
const redisAddress = "localhost:6379"

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

	// specifikace doby platnosti hodnoty uložené do Redisu
	expiration, err := time.ParseDuration("10s")
	if err != nil {
		panic(err)
	}

	// zápis hodnoty do databáze Redisu
	err = client.Set(context,
		"Seriál o jazyku Go",
		"https://www.root.cz/serialy/programovaci-jazyk-go/",
		expiration).Err()
	if err != nil {
		panic(err)
	}

	for {
		// přečtení hodnoty z databáze Redisu
		address, err := client.Get(context, "Seriál o jazyku Go").Result()

		// vyhodnocení předchozí operace
		switch {
		case err == redis.Nil:
			fmt.Println("no value found")
			return
		case err != nil:
			panic(err)
		default:
			fmt.Println("Adresa:", address)
		}

		time.Sleep(1 * time.Second)
	}
}
