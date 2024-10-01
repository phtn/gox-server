package config

import (
	"os"
)

type Config struct {
	ServerAddress string
}

func LoadConfig() Config {

	addr, exists := os.LookupEnv("ADDR")

	if !exists {
		addr = ":1981"
	}

	return Config{ServerAddress: addr}
}
