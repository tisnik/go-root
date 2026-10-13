// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá část
//    Zpracování konfiguračních souborů v Go s využitím knihovny Viper
//    https://www.root.cz/clanky/zpracovani-konfiguracnich-souboru-v-go-s-vyuzitim-knihovny-viper/
//
// Demonstrační příklad číslo 15:
//    Načtení konfigurace z proměnných prostředí.
//    Automatické navázání proměnné prostředí na konfigurační parametr.
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze čtyřicáté části:
//    https://github.com/tisnik/go-root/blob/master/article_40/README.md
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/html-root/article_40/15_env_variables_autobind.html
//

package main

import (
	"log"

	"github.com/spf13/viper"
)

func main() {
	log.Println("Reading configuration")

	viper.SetConfigName("config9")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("XTERM")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error in config file: %s \n", err)
	}
	log.Println("Done")

	log.Printf("Xterm locale: %s\n", viper.GetString("locale"))
	log.Printf("Xterm shell: %s\n", viper.GetString("shell"))
}
