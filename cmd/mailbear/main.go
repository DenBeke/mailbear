package main

import (
	"os"

	"github.com/DenBeke/mailbear"
	log "github.com/sirupsen/logrus"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {

	configFile := getenv("CONFIG_FILE", "config.yml")
	config, err := mailbear.GetConfigFromFile(configFile)
	if err != nil {
		log.Fatalf("couldn't read config: %v", err)
	}

	mailbear.Serve(config)

	return
}
