package config

// LoadGatewayConfig loads and decodes the "gateway." section of config.json
func LoadGatewayConfig() GatewayConfig {
	return loadConfig[GatewayConfig]("gateway")
}

// LoadIdentityConfig loads and decodes the "identity." section of config.json
func LoadIdentityConfig() IdentityConfig {
	return loadConfig[IdentityConfig]("identity")
}
