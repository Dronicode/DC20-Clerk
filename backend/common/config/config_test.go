package config_test

import (
	"os"
	"testing"

	"dc20clerk/common/config"

	"github.com/stretchr/testify/assert"
)

func writeConfig(t *testing.T, content string) string {
	t.Helper()

	var path string
	if _, err := os.Stat("/app/config.json"); err == nil {
		path = "/app/config.json"
	} else {
		path = "../config.json"
	}

	err := os.WriteFile(path, []byte(content), 0644)
	assert.NoError(t, err)
	return path
}

func TestLoadGatewayConfig_Valid(t *testing.T) {
	writeConfig(t, `{
        "gateway.IDENTITY_URL": "http://identity-service:8081"
    }`)

	cfg := config.LoadGatewayConfig()
	assert.Equal(t, "http://identity-service:8081", cfg.IdentityURL)
}

func TestLoadGatewayConfig_Malformed(t *testing.T) {
	writeConfig(t, `{ "gateway.IDENTITY_URL": "http://identity-service"`)

	assert.Panics(t, func() {
		_ = config.LoadGatewayConfig()
	})
}
