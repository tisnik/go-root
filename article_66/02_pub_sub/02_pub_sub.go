// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Šedesátá šestá část
//    Použití databáze Redis v aplikacích naprogramovaných v Go (2)
//    https://www.root.cz/clanky/pouziti-databaze-redis-v-aplikacich-naprogramovanych-v-go-2/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z šedesáté šesté části:
//    https://github.com/tisnik/go-root/blob/master/article_66/README.md

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// adresa určující službu Redisu, která se má použít
const redisAddress = "localhost:6379"

// jméno kanálu
const channelName = "c1"

// struktura obsahující "session"
type redisSession struct {
	client  *redis.Client
	context context.Context
}

// zdroj zpráv
func publisher(session redisSession, channel string, from int, to int) {
	for i := from; i < to; i++ {
		err := session.client.Publish(session.context, channel, i).Err()
		if err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
}

// příjemce zpráv
func subscriber(session redisSession, channel string) {
	pubsub := session.client.Subscribe(session.context, channel)
	for {
		message, err := pubsub.ReceiveMessage(session.context)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Channel: %s  Message: '%s'\n",
			message.Channel,
			message.Payload)
	}
}

// vstupní bod do demonstračního příkladu
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

	// vytvoříme session
	session := redisSession{
		client:  client,
		context: context,
	}

	// pokus o klasický handshake typu PING-PONG
	_, err := client.Ping(context).Result()
	if err != nil {
		panic(err)
	}

	// smazání kanálu, pokud existoval
	client.Del(context, channelName)

	// spustíme zdroje zpráv
	go publisher(session, channelName, 0, 10)

	// nyní můžeme spustit příjemce zpráv
	subscriber(session, channelName)
}
