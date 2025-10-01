package config

import (
	"os"
	"testing"
)

func TestLoadConfig_CeloSepoliaConfiguration(t *testing.T) {
	// Save original environment variables
	originalEnvVars := map[string]string{
		"CELO_SEPOLIA_RPC_URL":          os.Getenv("CELO_SEPOLIA_RPC_URL"),
		"CELO_SEPOLIA_CHAIN_ID":         os.Getenv("CELO_SEPOLIA_CHAIN_ID"),
		"CELO_SEPOLIA_CONTRACT_ADDRESS": os.Getenv("CELO_SEPOLIA_CONTRACT_ADDRESS"),
		"CELO_SEPOLIA_PRIVATE_KEY":      os.Getenv("CELO_SEPOLIA_PRIVATE_KEY"),
		"CELO_SEPOLIA_USDC_ADDRESS":     os.Getenv("CELO_SEPOLIA_USDC_ADDRESS"),
		"CONTRACT_ADDRESS":              os.Getenv("CONTRACT_ADDRESS"),
		"MERCHANT_PRIVATE_KEY":          os.Getenv("MERCHANT_PRIVATE_KEY"),
	}

	// Restore environment variables after test
	defer func() {
		for key, value := range originalEnvVars {
			if value == "" {
				os.Unsetenv(key)
			} else {
				os.Setenv(key, value)
			}
		}
	}()

	// Set required fields to prevent fatal errors
	os.Setenv("CONTRACT_ADDRESS", "0x123")
	os.Setenv("MERCHANT_PRIVATE_KEY", "0x456")

	tests := []struct {
		name     string
		envVars  map[string]string
		expected Config
	}{
		{
			name: "All Celo Sepolia environment variables set",
			envVars: map[string]string{
				"CELO_SEPOLIA_RPC_URL":          "https://alfajores-forno.celo-testnet.org",
				"CELO_SEPOLIA_CHAIN_ID":         "44787",
				"CELO_SEPOLIA_CONTRACT_ADDRESS": "0x1234567890123456789012345678901234567890",
				"CELO_SEPOLIA_PRIVATE_KEY":      "0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890",
				"CELO_SEPOLIA_USDC_ADDRESS":     "0x2F25deB3848C207fc8E0c34035B3Ba7fC157602B",
			},
			expected: Config{
				CeloSepoliaRPCURL:          "https://alfajores-forno.celo-testnet.org",
				CeloSepoliaChainID:         44787,
				CeloSepoliaContractAddress: "0x1234567890123456789012345678901234567890",
				CeloSepoliaPrivateKey:      "0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890",
				CeloSepoliaUSDCAddress:     "0x2F25deB3848C207fc8E0c34035B3Ba7fC157602B",
			},
		},
		{
			name: "Default Celo Sepolia chain ID when not specified",
			envVars: map[string]string{
				"CELO_SEPOLIA_RPC_URL":          "https://alfajores-forno.celo-testnet.org",
				"CELO_SEPOLIA_CONTRACT_ADDRESS": "0x1234567890123456789012345678901234567890",
				"CELO_SEPOLIA_PRIVATE_KEY":      "0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890",
				"CELO_SEPOLIA_USDC_ADDRESS":     "0x2F25deB3848C207fc8E0c34035B3Ba7fC157602B",
			},
			expected: Config{
				CeloSepoliaRPCURL:          "https://alfajores-forno.celo-testnet.org",
				CeloSepoliaChainID:         44787, // Default value
				CeloSepoliaContractAddress: "0x1234567890123456789012345678901234567890",
				CeloSepoliaPrivateKey:      "0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890",
				CeloSepoliaUSDCAddress:     "0x2F25deB3848C207fc8E0c34035B3Ba7fC157602B",
			},
		},
		{
			name: "Empty Celo Sepolia configuration",
			envVars: map[string]string{
				"CELO_SEPOLIA_RPC_URL":          "",
				"CELO_SEPOLIA_CHAIN_ID":         "",
				"CELO_SEPOLIA_CONTRACT_ADDRESS": "",
				"CELO_SEPOLIA_PRIVATE_KEY":      "",
				"CELO_SEPOLIA_USDC_ADDRESS":     "",
			},
			expected: Config{
				CeloSepoliaRPCURL:          "",
				CeloSepoliaChainID:         44787, // Default value
				CeloSepoliaContractAddress: "",
				CeloSepoliaPrivateKey:      "",
				CeloSepoliaUSDCAddress:     "",
			},
		},
		{
			name: "Invalid chain ID falls back to default",
			envVars: map[string]string{
				"CELO_SEPOLIA_RPC_URL":          "https://alfajores-forno.celo-testnet.org",
				"CELO_SEPOLIA_CHAIN_ID":         "invalid",
				"CELO_SEPOLIA_CONTRACT_ADDRESS": "0x1234567890123456789012345678901234567890",
				"CELO_SEPOLIA_PRIVATE_KEY":      "0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890",
				"CELO_SEPOLIA_USDC_ADDRESS":     "0x2F25deB3848C207fc8E0c34035B3Ba7fC157602B",
			},
			expected: Config{
				CeloSepoliaRPCURL:          "https://alfajores-forno.celo-testnet.org",
				CeloSepoliaChainID:         44787, // Default value when parsing fails
				CeloSepoliaContractAddress: "0x1234567890123456789012345678901234567890",
				CeloSepoliaPrivateKey:      "0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890",
				CeloSepoliaUSDCAddress:     "0x2F25deB3848C207fc8E0c34035B3Ba7fC157602B",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clear all Celo Sepolia environment variables first
			os.Unsetenv("CELO_SEPOLIA_RPC_URL")
			os.Unsetenv("CELO_SEPOLIA_CHAIN_ID")
			os.Unsetenv("CELO_SEPOLIA_CONTRACT_ADDRESS")
			os.Unsetenv("CELO_SEPOLIA_PRIVATE_KEY")
			os.Unsetenv("CELO_SEPOLIA_USDC_ADDRESS")

			// Set test environment variables
			for key, value := range tt.envVars {
				if value != "" {
					os.Setenv(key, value)
				}
			}

			config := LoadConfig()

			// Verify Celo Sepolia specific fields
			if config.CeloSepoliaRPCURL != tt.expected.CeloSepoliaRPCURL {
				t.Errorf("Expected CeloSepoliaRPCURL %s, got %s", tt.expected.CeloSepoliaRPCURL, config.CeloSepoliaRPCURL)
			}
			if config.CeloSepoliaChainID != tt.expected.CeloSepoliaChainID {
				t.Errorf("Expected CeloSepoliaChainID %d, got %d", tt.expected.CeloSepoliaChainID, config.CeloSepoliaChainID)
			}
			if config.CeloSepoliaContractAddress != tt.expected.CeloSepoliaContractAddress {
				t.Errorf("Expected CeloSepoliaContractAddress %s, got %s", tt.expected.CeloSepoliaContractAddress, config.CeloSepoliaContractAddress)
			}
			if config.CeloSepoliaPrivateKey != tt.expected.CeloSepoliaPrivateKey {
				t.Errorf("Expected CeloSepoliaPrivateKey %s, got %s", tt.expected.CeloSepoliaPrivateKey, config.CeloSepoliaPrivateKey)
			}
			if config.CeloSepoliaUSDCAddress != tt.expected.CeloSepoliaUSDCAddress {
				t.Errorf("Expected CeloSepoliaUSDCAddress %s, got %s", tt.expected.CeloSepoliaUSDCAddress, config.CeloSepoliaUSDCAddress)
			}
		})
	}
}

func TestLoadConfig_BackwardCompatibility(t *testing.T) {
	// Save original environment variables
	originalEnvVars := map[string]string{
		"CONTRACT_ADDRESS":     os.Getenv("CONTRACT_ADDRESS"),
		"MERCHANT_PRIVATE_KEY": os.Getenv("MERCHANT_PRIVATE_KEY"),
		"APTOS_NETWORK":        os.Getenv("APTOS_NETWORK"),
		"ETH_SEPOLIA_RPC_URL":  os.Getenv("ETH_SEPOLIA_RPC_URL"),
	}

	// Restore environment variables after test
	defer func() {
		for key, value := range originalEnvVars {
			if value == "" {
				os.Unsetenv(key)
			} else {
				os.Setenv(key, value)
			}
		}
	}()

	// Set required fields
	os.Setenv("CONTRACT_ADDRESS", "0x123")
	os.Setenv("MERCHANT_PRIVATE_KEY", "0x456")

	// Test that existing configuration still works
	os.Setenv("APTOS_NETWORK", "testnet")
	os.Setenv("ETH_SEPOLIA_RPC_URL", "https://sepolia.infura.io/v3/test")

	config := LoadConfig()

	// Verify existing fields are not affected by Celo Sepolia additions
	if config.AptosNetwork != "testnet" {
		t.Errorf("Expected AptosNetwork testnet, got %s", config.AptosNetwork)
	}
	if config.EVMRPCURL != "https://sepolia.infura.io/v3/test" {
		t.Errorf("Expected EVMRPCURL https://sepolia.infura.io/v3/test, got %s", config.EVMRPCURL)
	}

	// Verify Celo Sepolia fields have default values when not configured
	if config.CeloSepoliaChainID != 44787 {
		t.Errorf("Expected default CeloSepoliaChainID 44787, got %d", config.CeloSepoliaChainID)
	}
}

func TestGetEnvUint64_CeloSepoliaChainID(t *testing.T) {
	tests := []struct {
		name         string
		envValue     string
		defaultValue uint64
		expected     uint64
	}{
		{
			name:         "Valid chain ID",
			envValue:     "44787",
			defaultValue: 0,
			expected:     44787,
		},
		{
			name:         "Invalid chain ID returns default",
			envValue:     "invalid",
			defaultValue: 44787,
			expected:     44787,
		},
		{
			name:         "Empty chain ID returns default",
			envValue:     "",
			defaultValue: 44787,
			expected:     44787,
		},
		{
			name:         "Zero chain ID",
			envValue:     "0",
			defaultValue: 44787,
			expected:     0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original value
			original := os.Getenv("TEST_CHAIN_ID")
			defer func() {
				if original == "" {
					os.Unsetenv("TEST_CHAIN_ID")
				} else {
					os.Setenv("TEST_CHAIN_ID", original)
				}
			}()

			// Set test value
			if tt.envValue != "" {
				os.Setenv("TEST_CHAIN_ID", tt.envValue)
			} else {
				os.Unsetenv("TEST_CHAIN_ID")
			}

			result := getEnvUint64("TEST_CHAIN_ID", tt.defaultValue)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}
