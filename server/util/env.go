package util

import (
	"log"
	"os"
	. "strconv"

	env "github.com/joho/godotenv"
)

type EnvConfig struct {
	Port string
	ConnectionLimit int
	KeepAliveInterval int
}

var Config EnvConfig

func init() {
	err := env.Load()
	if err != nil {
		log.Fatalf("Failed to load env config: %s", err.Error())
	}

	Config = EnvConfig{
		Port: EnvGetString("PORT", true),
		ConnectionLimit: EnvGetNumber("CONNECTION_LIMIT", true),
		KeepAliveInterval: EnvGetNumber("KEEP_ALIVE_INTERVAL", true),
	}
}

// Errors if does not exist
func checkExists(key string, value string) {
	if len(value) == 0 {
		log.Fatalf("must provide '%s'", key)
	}
}

func EnvGetNumber(key string, required bool) int {
	valueRaw := os.Getenv(key)
	if required { checkExists(key, valueRaw) }

	result, err := Atoi(valueRaw)
	if err != nil {
		log.Fatal("port must be a number")
	}

	return result
}

func EnvGetString(key string, required bool) string {
	value := os.Getenv(key)
	if required { checkExists(key, value) }

	return value
}