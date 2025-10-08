// loader.go provides generic config loading logic.
// It reads config.json from a fixed path, filters by prefix, and decodes into typed structs.
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// LoadConfig reads config.json and filters keys by prefix (e.g. "gateway.")
func loadConfig[T any](prefix string) T {
	fmt.Println("[CONFIG] Loading config")
	path := resolveConfigPath()

	raw, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("[CONFIG] ✖ Failed to read config file: %w", err))
	}

	var all map[string]string
	if err := json.Unmarshal(raw, &all); err != nil {
		panic(fmt.Errorf("[CONFIG] ✖ Failed to parse config: %w", err))
	}

	filtered := make(map[string]string)
	for k, v := range all {
		if strings.HasPrefix(k, prefix) {
			filtered[strings.TrimPrefix(k, prefix+".")] = v
		}
	}

	// Marshal filtered map back to JSON and decode into typed struct
	buf, _ := json.Marshal(filtered)
	var cfg T
	if err := json.Unmarshal(buf, &cfg); err != nil {
		panic(fmt.Errorf("[CONFIG] ✖ Failed to decode typed config: %w", err))
	}

	return cfg
}

func resolveConfigPath() string {
	// Prefer container path
	if _, err := os.Stat("/app/config.json"); err == nil {
		fmt.Println("[CONFIG] Loading config from /app/config.json")
		return "/app/config.json"
	}

	// Fallback for local dev (from gateway/, identity/, etc.)
	if _, err := os.Stat("../config.json"); err == nil {
		fmt.Println("[CONFIG] Loading config from ../config.json")
		return "../config.json"
	}

	panic("[Config] ✖ config.json not found in /app/ or ../")
}
