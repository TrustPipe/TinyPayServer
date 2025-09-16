package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	// Aptos Network Configuration
	AptosNetwork   string
	AptosNodeURL   string
	AptosFaucetURL string

	// Contract Configuration
	ContractAddress string

	// Server Configuration
	Port string

	// Private Keys
	MerchantPrivateKey  string
	PaymasterPrivateKey string

	// Gas Configuration
	MaxGasAmount  uint64
	GasUnitPrice  uint64
}

func LoadConfig() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config := &Config{
		AptosNetwork:        getEnv("APTOS_NETWORK", "devnet"),
		AptosNodeURL:        getEnv("APTOS_NODE_URL", "https://fullnode.devnet.aptoslabs.com/v1"),
		AptosFaucetURL:      getEnv("APTOS_FAUCET_URL", "https://faucet.devnet.aptoslabs.com"),
		ContractAddress:     getEnv("CONTRACT_ADDRESS", ""),
		Port:                getEnv("PORT", "9090"),
		MerchantPrivateKey:  getEnv("MERCHANT_PRIVATE_KEY", ""),
		PaymasterPrivateKey: getEnv("PAYMASTER_PRIVATE_KEY", ""),
		MaxGasAmount:        getEnvUint64("MAX_GAS_AMOUNT", 100000),
		GasUnitPrice:        getEnvUint64("GAS_UNIT_PRICE", 100),
	}

	// Validate required fields
	if config.ContractAddress == "" {
		log.Fatal("CONTRACT_ADDRESS is required")
	}
	if config.MerchantPrivateKey == "" {
		log.Fatal("MERCHANT_PRIVATE_KEY is required")
	}

	return config
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvUint64(key string, defaultValue uint64) uint64 {
	if value := os.Getenv(key); value != "" {
		if parsed, err := strconv.ParseUint(value, 10, 64); err == nil {
			return parsed
		}
	}
	return defaultValue
}
