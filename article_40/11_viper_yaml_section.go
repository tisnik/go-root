package main

import (
	"github.com/spf13/viper"
	"log"
)

func main() {
	log.Println("Reading configuration")

	viper.SetConfigName("config9")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error in config file: %s \n", err)
	}
	log.Println("Done")

	serviceConfig := viper.Sub("service")
	url := serviceConfig.GetString("url")
	port := serviceConfig.GetInt("port")

	log.Printf("Starting the service at address %s:%d\n", url, port)
}
