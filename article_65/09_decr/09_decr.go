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

	// pokus o klasický handshake typu PING-PONG
	_, err := client.Ping(context).Result()
	if err != nil {
		panic(err)
	}

	// smazání hodnoty, pokud existovala
	client.Del(context, "counter")

	// dekrementace (neexistující) hodnoty
	newValue := client.Decr(context, "counter").Val()
	fmt.Println("Counter value:", newValue)

	// přečtení hodnoty z databáze Redisu
	newValue = client.DecrBy(context, "counter", 0).Val()
	fmt.Println("Counter value:", newValue)

	// dekrementace (nyní již existující) hodnoty
	newValue = client.Decr(context, "counter").Val()
	fmt.Println("Counter value:", newValue)

	// inkrementace (nyní již existující) hodnoty
	newValue = client.DecrBy(context, "counter", -1).Val()
	fmt.Println("Counter value:", newValue)
}
