// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá část
//    Zpracování konfiguračních souborů v Go s využitím knihovny Viper
//    https://www.root.cz/clanky/zpracovani-konfiguracnich-souboru-v-go-s-vyuzitim-knihovny-viper/
//
// Demonstrační příklad číslo 5:
//    Přečtení konfigurace obsahující pole.
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze čtyřicáté části:
//    https://github.com/tisnik/go-root/blob/master/article_40/README.md
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/html-root/article_40/05_toml_array.html
//

package main

import (
	"log"

	"github.com/spf13/viper"
)

func main() {
	log.Println("Reading configuration")

	viper.SetConfigName("config3")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error in config file: %s \n", err)
	}
	log.Println("Done")

	serviceConfig := viper.Sub("service")
	url := serviceConfig.GetString("url")
	port := serviceConfig.GetInt("port")

	usersConfig := viper.Sub("users")
	accepted := usersConfig.GetStringSlice("accepted")
	blacklisted := usersConfig.GetStringSlice("blacklisted")

	log.Printf("Starting the service at address %s:%d\n", url, port)
	log.Printf("Accepted users: %v\n", accepted)
	log.Printf("Blacklisted users: %v\n", blacklisted)
}
