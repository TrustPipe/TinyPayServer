package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/pelletier/go-toml/v2"
)

// EVMToken represents an ERC20 token configuration
type EVMToken struct {
	Symbol  string `toml:"symbol"`
	Address string `toml:"address"`
}

// EVMNativeToken represents the native token configuration
type EVMNativeToken struct {
	Symbol  string `toml:"symbol"`
	Address string `toml:"address"`
}

// EVMNetwork represents a single EVM network configuration
type EVMNetwork struct {
	Name         string         `toml:"name"`
	RPCURL       string         `toml:"rpc_url"`
	ChainID      uint64         `toml:"chain_id"`
	ContractAddress string      `toml:"contract_address"`
	PrivateKey   string         `toml:"private_key"`
	NativeToken  EVMNativeToken `toml:"native_token"`
	Tokens       []EVMToken     `toml:"tokens"`
}

// TomlConfig represents the TOML configuration structure
type TomlConfig struct {
	Aptos struct {
		Network   string `toml:"network"`
		NodeURL   string `toml:"node_url"`
		FaucetURL string `toml:"faucet_url"`
	} `toml:"aptos"`
	
	Contract struct {
		Address           string `toml:"address"`
		USDCMetadataAddress string `toml:"usdc_metadata_address"`
	} `toml:"contract"`
	
	Server struct {
		Port string `toml:"port"`
	} `toml:"server"`
	
	Gas struct {
		MaxGasAmount uint64 `toml:"max_gas_amount"`
		GasUnitPrice uint64 `toml:"gas_unit_price"`
	} `toml:"gas"`
	
	Keys struct {
		MerchantPrivateKey  string `toml:"merchant_private_key"`
		PaymasterPrivateKey string `toml:"paymaster_private_key"`
	} `toml:"keys"`
	
	EVMNetworks []EVMNetwork `toml:"evm_networks"`
}

// Config represents the application configuration (legacy structure for compatibility)
type Config struct {
	// Aptos Network Configuration
	AptosNetwork   string
	AptosNodeURL   string
	AptosFaucetURL string

	// Contract Configuration
	ContractAddress     string
	USDCContractAddress string // Legacy coin type address
	USDCMetadataAddress string // FA metadata address

	// EVM Configuration (legacy fields for eth-sepolia)
	EVMRPCURL             string
	EVMChainID            uint64
	EVMContractAddress    string
	EVMPrivateKey         string
	ETHSepoliaUSDCAddress string // Test USDC token address for EVM

	// Celo Sepolia Configuration
	CeloSepoliaRPCURL          string
	CeloSepoliaChainID         uint64
	CeloSepoliaContractAddress string
	CeloSepoliaPrivateKey      string
	CeloSepoliaUSDCAddress     string

	// Server Configuration
	Port string

	// Private Keys
	MerchantPrivateKey  string
	PaymasterPrivateKey string

	// Gas Configuration
	MaxGasAmount uint64
	GasUnitPrice uint64
	
	// EVM Networks (new array-based configuration)
	EVMNetworks []EVMNetwork
}

func LoadConfig() *Config {
	// Try to load TOML config first
	config := loadTomlConfig()
	if config != nil {
		return config
	}
	
	// Fallback to legacy .env loading
	return loadEnvConfig()
}

func loadTomlConfig() *Config {
	// Try to load config.toml file
	tomlData, err := os.ReadFile("config.toml")
	if err != nil {
		return nil // TOML file not found, fallback to .env
	}
	
	var tomlConfig TomlConfig
	if err := toml.Unmarshal(tomlData, &tomlConfig); err != nil {
		log.Printf("Error parsing TOML config: %v, falling back to .env", err)
		return nil
	}
	
	// Convert TOML config to legacy Config struct
	config := &Config{
		// Aptos configuration
		AptosNetwork:        tomlConfig.Aptos.Network,
		AptosNodeURL:        tomlConfig.Aptos.NodeURL,
		AptosFaucetURL:      tomlConfig.Aptos.FaucetURL,
		
		// Contract configuration
		ContractAddress:       tomlConfig.Contract.Address,
		USDCMetadataAddress:   tomlConfig.Contract.USDCMetadataAddress,
		
		// Server configuration
		Port:                  tomlConfig.Server.Port,
		
		// Gas configuration
		MaxGasAmount:          tomlConfig.Gas.MaxGasAmount,
		GasUnitPrice:          tomlConfig.Gas.GasUnitPrice,
		
		// Private keys
		MerchantPrivateKey:    tomlConfig.Keys.MerchantPrivateKey,
		PaymasterPrivateKey:   tomlConfig.Keys.PaymasterPrivateKey,
		
		// EVM networks (new array-based configuration)
		EVMNetworks:           tomlConfig.EVMNetworks,
	}
	
    // Legacy EVM fields are intentionally not set to avoid hardcoding network names.
	
	// Validate required fields
	if config.ContractAddress == "" {
		log.Fatal("Contract address is required in TOML config")
	}
	if config.MerchantPrivateKey == "" {
		log.Fatal("Merchant private key is required in TOML config")
	}
	
	return config
}

func loadEnvConfig() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config := &Config{
		AptosNetwork:               getEnv("APTOS_NETWORK", "devnet"),
		AptosNodeURL:               getEnv("APTOS_NODE_URL", "https://fullnode.devnet.aptoslabs.com/v1"),
		AptosFaucetURL:             getEnv("APTOS_FAUCET_URL", "https://faucet.devnet.aptoslabs.com"),
		ContractAddress:            getEnv("CONTRACT_ADDRESS", ""),
		USDCContractAddress:        getEnv("USDC_CONTRACT_ADDRESS", "0xaadbf0681ef3dc9decd123340db16954f85319853533ed4ace6ec5d11aaad190::test_usdc::TestUSDC"),
		USDCMetadataAddress:        getEnv("USDC_METADATA_ADDRESS", "0x331ebb81b96e2b0114a68a070d433ac9659361f1eab45f831a437df1fde51fde"),
		EVMRPCURL:                  getEnv("ETH_SEPOLIA_RPC_URL", ""),
		EVMChainID:                 getEnvUint64("ETH_SEPOLIA_CHAIN_ID", 0),
		EVMContractAddress:         getEnv("ETH_SEPOLIA_CONTRACT_ADDRESS", ""),
		EVMPrivateKey:              getEnv("ETH_SEPOLIA_PRIVATE_KEY", ""),
		ETHSepoliaUSDCAddress:      getEnv("ETH_SEPOLIA_USDC_ADDRESS", "0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238"), // Default to provided test address
		CeloSepoliaRPCURL:          getEnv("CELO_SEPOLIA_RPC_URL", ""),
		CeloSepoliaChainID:         getEnvUint64("CELO_SEPOLIA_CHAIN_ID", 44787),
		CeloSepoliaContractAddress: getEnv("CELO_SEPOLIA_CONTRACT_ADDRESS", ""),
		CeloSepoliaPrivateKey:      getEnv("CELO_SEPOLIA_PRIVATE_KEY", ""),
		CeloSepoliaUSDCAddress:     getEnv("CELO_SEPOLIA_USDC_ADDRESS", ""),
		Port:                       getEnv("PORT", "9090"),
		MerchantPrivateKey:         getEnv("MERCHANT_PRIVATE_KEY", ""),
		PaymasterPrivateKey:        getEnv("PAYMASTER_PRIVATE_KEY", ""),
		MaxGasAmount:               getEnvUint64("MAX_GAS_AMOUNT", 100000),
		GasUnitPrice:               getEnvUint64("GAS_UNIT_PRICE", 100),
	}
	
    // Skip env-to-array conversion to avoid embedding hardcoded network names.

	// Validate required fields
	if config.ContractAddress == "" {
		log.Fatal("CONTRACT_ADDRESS is required")
	}
	if config.MerchantPrivateKey == "" {
		log.Fatal("MERCHANT_PRIVATE_KEY is required")
	}

	// Validate Celo Sepolia configuration (log warnings but continue operation)
	if config.CeloSepoliaRPCURL == "" {
		log.Println("Warning: CELO_SEPOLIA_RPC_URL not configured, Celo Sepolia network will be unavailable")
	}
	if config.CeloSepoliaContractAddress == "" {
		log.Println("Warning: CELO_SEPOLIA_CONTRACT_ADDRESS not configured, Celo Sepolia payments will be unavailable")
	}
	if config.CeloSepoliaPrivateKey == "" {
		log.Println("Warning: CELO_SEPOLIA_PRIVATE_KEY not configured, Celo Sepolia payments will be unavailable")
	}
	if config.CeloSepoliaUSDCAddress == "" {
		log.Println("Warning: CELO_SEPOLIA_USDC_ADDRESS not configured, USDC payments on Celo Sepolia will be unavailable")
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
