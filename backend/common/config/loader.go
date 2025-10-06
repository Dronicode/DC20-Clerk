package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// LoadConfig reads config.json and filters keys by prefix (e.g. "gateway.")
func loadConfig(path string, prefix string) (map[string]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var raw map[string]string
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("invalid config format: %w", err)
	}

	filtered := make(map[string]string)
	for k, v := range raw {
		if strings.HasPrefix(k, prefix) {
			trimmed := strings.TrimPrefix(k, prefix)
			filtered[trimmed] = v
		}
	}

	if len(filtered) == 0 {
		return nil, fmt.Errorf("no config keys found for prefix: %s", prefix)
	}

	return filtered, nil
}
