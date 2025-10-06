package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Dronicode/DC20-Clerk/common/config"

	"github.com/stretchr/testify/assert"
)

func writeTempFile(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "config.json")
	err := os.WriteFile(path, []byte(content), 0644)
	assert.NoError(t, err)
	return path
}

func TestLoadConfig_ValidPrefix(t *testing.T) {
	content := `{
        "gateway.IDENTITY_URL": "http://identity-service:8081",
        "identity.SUPABASE_URL": "https://supabase.co"
    }`
	path := writeTempFile(t, content)

	cfg, err := config.LoadConfig(path, "gateway.")
	assert.NoError(t, err)
	assert.Equal(t, "http://identity-service:8081", cfg["IDENTITY_URL"])
}

func TestLoadConfig_NoMatchingPrefix(t *testing.T) {
	content := `{
        "identity.SUPABASE_URL": "https://supabase.co"
    }`
	path := writeTempFile(t, content)

	_, err := config.LoadConfig(path, "gateway.")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no config keys found")
}

func TestLoadConfig_MalformedJSON(t *testing.T) {
	content := `{ "gateway.IDENTITY_URL": "http://identity-service"`
	path := writeTempFile(t, content)

	_, err := config.LoadConfig(path, "gateway.")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid config format")
}

func TestLoadConfig_FileMissing(t *testing.T) {
	_, err := config.LoadConfig("nonexistent.json", "gateway.")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to read config file")
}
