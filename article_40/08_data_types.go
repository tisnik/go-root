// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá část
//    Zpracování konfiguračních souborů v Go s využitím knihovny Viper
//    https://www.root.cz/clanky/zpracovani-konfiguracnich-souboru-v-go-s-vyuzitim-knihovny-viper/

package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	log.Println("Reading configuration")

	viper.SetConfigName("config6")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error in config file: %s \n", err)
	}
	log.Println("Done")

	fmt.Printf("integer1: %d\n", viper.GetInt("integer1"))
	fmt.Printf("integer2: %d\n", viper.GetInt("integer2"))

	fmt.Printf("float1: %f\n", viper.GetFloat64("float1"))
	fmt.Printf("float2: %f\n", viper.GetFloat64("float2"))
	fmt.Printf("float3: %f\n", viper.GetFloat64("float3"))
	fmt.Printf("float4: %f\n", viper.GetFloat64("float4"))

	fmt.Printf("bool1: %t\n", viper.GetBool("bool1"))
	fmt.Printf("bool2: %t\n", viper.GetBool("bool2"))

	fmt.Printf("date1: %s\n", viper.GetTime("date1").Format("Mon Jan 2 15:04:05 MST 2006"))
	fmt.Printf("date2: %s\n", viper.GetTime("date2").Format("Mon Jan 2 15:04:05 MST 2006"))
	fmt.Printf("date3: %s\n", viper.GetTime("date3").Format("Mon Jan 2 15:04:05 MST 2006"))
	fmt.Printf("date4: %s\n", viper.GetTime("date4").Format("Mon Jan 2 15:04:05 MST 2006"))
}
