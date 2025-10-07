package config

type GatewayConfig struct {
	IdentityURL string `json:"IDENTITY_URL"`
}

type IdentityConfig struct {
	SupabaseURL       string `json:"SUPABASE_URL"`
	SupabaseSecretKey string `json:"SUPABASE_SECRET_KEY"`
}
