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
	ContractAddress     string
	USDCContractAddress string // Legacy coin type address
	USDCMetadataAddress string // FA metadata address

	// Server Configuration
	Port string

	// Private Keys
	MerchantPrivateKey  string
	PaymasterPrivateKey string

	// Gas Configuration
	MaxGasAmount uint64
	GasUnitPrice uint64
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
		USDCContractAddress: getEnv("USDC_CONTRACT_ADDRESS", "0xaadbf0681ef3dc9decd123340db16954f85319853533ed4ace6ec5d11aaad190::test_usdc::TestUSDC"),
		USDCMetadataAddress: getEnv("USDC_METADATA_ADDRESS", "0x331ebb81b96e2b0114a68a070d433ac9659361f1eab45f831a437df1fde51fde"),
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
