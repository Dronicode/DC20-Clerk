package util

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var loadOnce sync.Once

// TODO DEPRECATED in favor of configs, remove when they definitey fully work.
func Env(key string) string {
	loadOnce.Do(func() {
		if err := godotenv.Load(".env"); err != nil {
			log.Printf("Warning: .env failed to load: %v", err)
		}
	})

	val := os.Getenv(key)
	if val == "" {
		log.Printf("Warning: env var %s is not set", key)
	}
	return val
}
