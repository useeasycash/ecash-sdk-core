package config

import (
	"os"
	"time"
)

// SDKConfig holds global configuration for the SDK
type SDKConfig struct {
	// API Configuration
	APIEndpoint string
	APIKey      string
	Environment string // "mainnet" | "testnet" | "devnet"

	// Network Configuration
	Timeout      time.Duration
	MaxRetries   int
	RetryBackoff time.Duration

	// Privacy Configuration
	EnableZKProofs bool
	ProofCacheTTL  time.Duration

	// Performance Configuration
	EnableMetrics bool
	EnableCaching bool
	CacheTTL      time.Duration
}

// DefaultConfig returns sensible defaults
func DefaultConfig() *SDKConfig {
	return &SDKConfig{
		APIEndpoint:    getEnv("ECASH_API_ENDPOINT", "https://api.useeasy.cash"),
		APIKey:         getEnv("ECASH_API_KEY", ""),
		Environment:    getEnv("ECASH_ENV", "mainnet"),
		Timeout:        30 * time.Second,
		MaxRetries:     3,
		RetryBackoff:   2 * time.Second,
		EnableZKProofs: true,
		ProofCacheTTL:  5 * time.Minute,
		EnableMetrics:  true,
		EnableCaching:  true,
		CacheTTL:       1 * time.Minute,
	}
}

// getEnv retrieves environment variable with fallback
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// Validate checks if the configuration is valid
func (c *SDKConfig) Validate() error {
	// Add validation logic here
	return nil
}
