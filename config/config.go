package config

import (
	"log"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	JWT_SECRET  string
}

func GetConfig() Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf("config/config.json", &configuration)
	if err != nil {
		log.Fatalf("Failed to load configuration file, error: %v", err)
	}
	return configuration
}
