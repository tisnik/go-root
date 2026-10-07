// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá část
//    Zpracování konfiguračních souborů v Go s využitím knihovny Viper
//    https://www.root.cz/clanky/zpracovani-konfiguracnich-souboru-v-go-s-vyuzitim-knihovny-viper/
//
// Demonstrační příklad číslo 2:
//    Základní způsob čtení konfigurace ze souboru s formátem TOML.
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze čtyřicáté části:
//    https://github.com/tisnik/go-root/blob/master/article_40/README.md
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/html-root/article_40/02_viper_read_toml.html
//

package main

import (
	"log"

	"github.com/spf13/viper"
)

func main() {
	log.Println("Reading configuration")

	viper.SetConfigName("config1")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error in config file: %s \n", err)
	}
	log.Println("Done")

	url := viper.GetString("url")
	port := viper.GetInt("port")

	log.Printf("Starting the service at address %s:%d\n", url, port)
}
