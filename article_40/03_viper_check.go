// Seriál "Programovací jazyk Go"
//
// Čtyřicátá část
//    Zpracování konfiguračních souborů v Go s využitím knihovny Viper
//    https://www.root.cz/clanky/zpracovani-konfiguracnich-souboru-v-go-s-vyuzitim-knihovny-viper/

package main

import (
	"github.com/spf13/viper"
	"log"
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

	if !viper.IsSet("url") {
		log.Fatalf("URL is not specified in config file")
	}

	if !viper.IsSet("port") {
		log.Fatalf("port is not specified in config file")
	}

	url := viper.GetString("url")
	port := viper.GetInt("port")

	log.Printf("Starting the service at address %s:%d\n", url, port)
}
