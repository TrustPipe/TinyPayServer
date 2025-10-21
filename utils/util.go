package utils

import (
	"fmt"
	"strings"
	"tinypay-server/config"

	"github.com/gagliardetto/solana-go"
)

// HexToASCIIBytes 将十六进制字符串转换为 ASCII 字节数组
func HexToASCIIBytes(hexStr string) []byte {
	// 清理输入
	hexStr = strings.ToLower(strings.TrimSpace(hexStr))

	return []byte(strings.TrimPrefix(hexStr, "0x"))
}

// GetMetadataMapping 根据配置获取币种到 FA metadata 地址的映射
func GetMetadataMapping(cfg *config.Config) map[string]string {
	return map[string]string{
		"APT":  "0xa", // Normalize to lowercase for comparisons
		"USDC": strings.ToLower(cfg.USDCMetadataAddress),
	}
}

// GetCoinType 根据币种名称获取对应的合约类型
func GetCoinType(cfg *config.Config, currency string) (string, error) {
	if currency == "" {
		currency = "APT" // 默认为APT
	}

	coinTypeMapping := GetMetadataMapping(cfg)
	coinType, exists := coinTypeMapping[strings.ToUpper(currency)]
	if !exists {
		return "", fmt.Errorf("unsupported currency: %s", currency)
	}

	return strings.ToLower(coinType), nil
}

// GetSupportedCurrencies 获取支持的币种列表
func GetSupportedCurrencies(cfg *config.Config) []string {
	coinTypeMapping := GetMetadataMapping(cfg)
	currencies := make([]string, 0, len(coinTypeMapping))
	for currency := range coinTypeMapping {
		currencies = append(currencies, currency)
	}
	return currencies
}

// GetMetadataAddress 根据币种名称获取对应的 FA metadata 地址
func GetMetadataAddress(cfg *config.Config, currency string) (string, error) {
	if currency == "" {
		currency = "APT" // 默认为APT
	}

	metadataMapping := GetMetadataMapping(cfg)
	metadataAddr, exists := metadataMapping[strings.ToUpper(currency)]
	if !exists {
		return "", fmt.Errorf("unsupported currency: %s", currency)
	}

	return strings.ToLower(metadataAddr), nil
}

// GetCurrencyFromCoinType 根据合约类型获取币种名称（反向映射）
func GetCurrencyFromCoinType(cfg *config.Config, coinType string) string {
	coinTypeMapping := GetMetadataMapping(cfg)
	for currency, contractType := range coinTypeMapping {
		if strings.EqualFold(contractType, coinType) {
			return currency
		}
	}
	return "UNKNOWN"
}

// GetEVMNetworkConfig returns the EVM network configuration by name
func GetEVMNetworkConfig(cfg *config.Config, networkName string) *config.EVMNetwork {
	if cfg == nil {
		return nil
	}

	for _, network := range cfg.EVMNetworks {
		if strings.EqualFold(network.Name, networkName) {
			return &network
		}
	}

	return nil
}

// GetEVMTokenMappingByNetwork 获取EVM代币地址映射 (支持多网络)
func GetEVMTokenMappingByNetwork(cfg *config.Config, network string) map[string]string {
	// First try to get token mapping from the new EVMNetworks configuration
	for _, evmNetwork := range cfg.EVMNetworks {
		if strings.EqualFold(evmNetwork.Name, network) {
			tokenMapping := make(map[string]string)

			// Add native token
			tokenMapping[evmNetwork.NativeToken.Symbol] = strings.ToLower(evmNetwork.NativeToken.Address)

			// Add ERC20 tokens
			for _, token := range evmNetwork.Tokens {
				tokenMapping[token.Symbol] = strings.ToLower(token.Address)
			}

			return tokenMapping
		}
	}

	// If not found, return empty mapping (no hardcoded defaults)
	return map[string]string{}
}

// GetEVMTokenMapping 获取EVM代币地址映射 (向后兼容，默认以太坊Sepolia)
func GetEVMTokenMapping(cfg *config.Config) map[string]string {
	if cfg == nil || len(cfg.EVMNetworks) == 0 {
		return map[string]string{}
	}
	return GetEVMTokenMappingByNetwork(cfg, cfg.EVMNetworks[0].Name)
}

// GetEVMTokenAddressByNetwork 根据币种获取EVM代币地址 (支持多网络)
func GetEVMTokenAddressByNetwork(cfg *config.Config, currency string, network string) (string, error) {
	if currency == "" {
		if netCfg := GetEVMNetworkConfig(cfg, network); netCfg != nil {
			currency = netCfg.NativeToken.Symbol
		}
	}

	tokenMapping := GetEVMTokenMappingByNetwork(cfg, network)
	tokenAddress, exists := tokenMapping[strings.ToUpper(currency)]
	if !exists {
		return "", fmt.Errorf("unsupported EVM currency: %s for network: %s", currency, network)
	}

	return strings.ToLower(tokenAddress), nil
}

// GetEVMTokenAddress 根据币种获取EVM代币地址 (向后兼容，默认以太坊Sepolia)
func GetEVMTokenAddress(cfg *config.Config, currency string) (string, error) {
	if cfg == nil || len(cfg.EVMNetworks) == 0 {
		return "", fmt.Errorf("no EVM networks configured")
	}
	return GetEVMTokenAddressByNetwork(cfg, currency, cfg.EVMNetworks[0].Name)
}

// GetCurrencyFromEVMTokenAddressByNetwork 根据EVM代币地址获取币种名称 (支持多网络)
func GetCurrencyFromEVMTokenAddressByNetwork(cfg *config.Config, tokenAddress string, network string) string {
	target := strings.ToLower(strings.TrimSpace(tokenAddress))
	tokenMapping := GetEVMTokenMappingByNetwork(cfg, network)
	for currency, addr := range tokenMapping {
		if strings.ToLower(strings.TrimSpace(addr)) == target {
			return currency
		}
	}
	return "UNKNOWN"
}

// GetCurrencyFromEVMTokenAddress 根据EVM代币地址获取币种名称 (向后兼容，默认以太坊Sepolia)
func GetCurrencyFromEVMTokenAddress(cfg *config.Config, tokenAddress string) string {
	if cfg == nil || len(cfg.EVMNetworks) == 0 {
		return "UNKNOWN"
	}
	return GetCurrencyFromEVMTokenAddressByNetwork(cfg, tokenAddress, cfg.EVMNetworks[0].Name)
}

// GetCurrencyFromMetadata 根据 metadata 地址获取币种名称（反向映射）
func GetCurrencyFromMetadata(cfg *config.Config, metadataAddr string) string {
	target := strings.ToLower(strings.TrimSpace(metadataAddr))
	metadataMapping := GetMetadataMapping(cfg)
	for currency, addr := range metadataMapping {
		if strings.ToLower(strings.TrimSpace(addr)) == target {
			return currency
		}
	}
	return "UNKNOWN"
}

// Deprecated legacy Celo helpers removed in favor of dynamic configuration

// Deprecated legacy Celo helpers removed in favor of dynamic configuration

// Deprecated legacy Celo helpers removed in favor of dynamic configuration

// NetworkCurrencyValidationMatrix 定义网络-货币组合验证矩阵
type NetworkCurrencyValidationMatrix struct {
	validCombinations map[string][]string
}

// NewNetworkCurrencyValidationMatrix 创建新的网络-货币验证矩阵
func NewNetworkCurrencyValidationMatrix(cfg *config.Config) *NetworkCurrencyValidationMatrix {
	combos := map[string][]string{
		"aptos-testnet": {"APT", "USDC"},
	}
	if cfg != nil {
		// Add EVM networks
		for _, net := range cfg.EVMNetworks {
			currencies := []string{strings.ToUpper(net.NativeToken.Symbol)}
			for _, t := range net.Tokens {
				currencies = append(currencies, strings.ToUpper(t.Symbol))
			}
			combos[strings.ToLower(net.Name)] = currencies
		}
		
		// Add Solana networks
		for _, net := range cfg.SolanaNetworks {
			currencies := []string{strings.ToUpper(net.NativeToken.Symbol)}
			for _, t := range net.Tokens {
				currencies = append(currencies, strings.ToUpper(t.Symbol))
			}
			combos[strings.ToLower(net.Name)] = currencies
		}
	}
	return &NetworkCurrencyValidationMatrix{validCombinations: combos}
}

// IsValidCombination 检查网络-货币组合是否有效
func (m *NetworkCurrencyValidationMatrix) IsValidCombination(network, currency string) bool {
	if validCurrencies, networkExists := m.validCombinations[strings.ToLower(network)]; networkExists {
		for _, validCurrency := range validCurrencies {
			if strings.EqualFold(currency, validCurrency) {
				return true
			}
		}
	}
	return false
}

// GetSupportedCurrenciesForNetwork 获取指定网络支持的货币列表
func (m *NetworkCurrencyValidationMatrix) GetSupportedCurrenciesForNetwork(network string) []string {
	if validCurrencies, networkExists := m.validCombinations[strings.ToLower(network)]; networkExists {
		return validCurrencies
	}
	return []string{}
}

// GetSupportedNetworks 获取所有支持的网络列表
func (m *NetworkCurrencyValidationMatrix) GetSupportedNetworks() []string {
	networks := make([]string, 0, len(m.validCombinations))
	for network := range m.validCombinations {
		networks = append(networks, network)
	}
	return networks
}

// ValidateNetworkCurrencyCombination 验证网络-货币组合并返回详细错误信息
func ValidateNetworkCurrencyCombination(cfg *config.Config, network, currency string) error {
	matrix := NewNetworkCurrencyValidationMatrix(cfg)

	// 检查网络是否支持
	supportedNetworks := matrix.GetSupportedNetworks()
	networkSupported := false
	for _, supportedNet := range supportedNetworks {
		if strings.EqualFold(network, supportedNet) {
			networkSupported = true
			break
		}
	}

	if !networkSupported {
		return fmt.Errorf("unsupported network '%s'. Supported networks: %s",
			network, strings.Join(supportedNetworks, ", "))
	}

	// 检查货币是否在该网络上支持
	if !matrix.IsValidCombination(network, currency) {
		supportedCurrencies := matrix.GetSupportedCurrenciesForNetwork(network)
		return fmt.Errorf("currency '%s' is not supported on network '%s'. Supported currencies for %s: %s",
			currency, network, network, strings.Join(supportedCurrencies, ", "))
	}

	return nil
}

// GetDefaultCurrencyForNetwork 获取网络的默认货币
func GetDefaultCurrencyForNetwork(cfg *config.Config, network string) string {
	if strings.ToLower(network) == "aptos-testnet" {
		return "APT"
	}
	if netCfg := GetEVMNetworkConfig(cfg, network); netCfg != nil {
		return strings.ToUpper(netCfg.NativeToken.Symbol)
	}
	if netCfg := GetSolanaNetworkConfig(cfg, network); netCfg != nil {
		return strings.ToUpper(netCfg.NativeToken.Symbol)
	}
	return ""
}

// IsNativeCurrency 检查货币是否为指定网络的原生货币
func IsNativeCurrency(cfg *config.Config, network, currency string) bool {
	if strings.ToLower(network) == "aptos-testnet" {
		return strings.EqualFold(currency, "APT")
	}
	if netCfg := GetEVMNetworkConfig(cfg, network); netCfg != nil {
		return strings.EqualFold(strings.ToUpper(currency), strings.ToUpper(netCfg.NativeToken.Symbol))
	}
	if netCfg := GetSolanaNetworkConfig(cfg, network); netCfg != nil {
		return strings.EqualFold(strings.ToUpper(currency), strings.ToUpper(netCfg.NativeToken.Symbol))
	}
	return false
}

// ValidateTokenAddressForNetworkCurrency 验证代币地址是否与网络-货币组合匹配
func ValidateTokenAddressForNetworkCurrency(cfg *config.Config, network, currency, tokenAddress string) error {
	// 首先验证网络-货币组合是否有效
	if err := ValidateNetworkCurrencyCombination(cfg, network, currency); err != nil {
		return err
	}

	// 获取预期的代币地址
	var expectedAddress string
	var err error

	switch strings.ToLower(network) {
	case "aptos-testnet":
		// Aptos 使用 metadata 地址而不是代币地址
		expectedAddress, err = GetMetadataAddress(cfg, currency)
	default:
		// Any configured EVM network
		expectedAddress, err = GetEVMTokenAddressByNetwork(cfg, currency, network)
	}

	if err != nil {
		return fmt.Errorf("failed to get expected token address for %s on %s: %w", currency, network, err)
	}

	// 比较地址（不区分大小写）
	if !strings.EqualFold(strings.TrimSpace(tokenAddress), strings.TrimSpace(expectedAddress)) {
		return fmt.Errorf("token address mismatch for %s on %s: expected %s, got %s",
			currency, network, expectedAddress, tokenAddress)
	}

	return nil
}

// ParseSolanaPublicKey parses a Solana public key from base58 string
func ParseSolanaPublicKey(address string) (solana.PublicKey, error) {
	pubkey, err := solana.PublicKeyFromBase58(address)
	if err != nil {
		return solana.PublicKey{}, fmt.Errorf("invalid Solana address: %w", err)
	}
	return pubkey, nil
}

// GetSolanaNetworkConfig returns the Solana network configuration by name
func GetSolanaNetworkConfig(cfg *config.Config, networkName string) *config.SolanaNetwork {
	if cfg == nil {
		return nil
	}

	for i := range cfg.SolanaNetworks {
		if strings.EqualFold(cfg.SolanaNetworks[i].Name, networkName) {
			return &cfg.SolanaNetworks[i]
		}
	}

	return nil
}

// GetSolanaTokenMappingByNetwork 获取Solana代币地址映射 (支持多网络)
func GetSolanaTokenMappingByNetwork(cfg *config.Config, network string) map[string]string {
	for _, solanaNetwork := range cfg.SolanaNetworks {
		if strings.EqualFold(solanaNetwork.Name, network) {
			tokenMapping := make(map[string]string)

			// Add native token (SOL)
			tokenMapping[solanaNetwork.NativeToken.Symbol] = ""

			// Add SPL tokens
			for _, token := range solanaNetwork.Tokens {
				tokenMapping[strings.ToUpper(token.Symbol)] = strings.ToLower(token.Address)
			}

			return tokenMapping
		}
	}

	return map[string]string{}
}

// GetSolanaTokenAddressByNetwork 获取特定网络的代币地址
func GetSolanaTokenAddressByNetwork(cfg *config.Config, currency, network string) (string, error) {
	tokenMapping := GetSolanaTokenMappingByNetwork(cfg, network)
	address, exists := tokenMapping[strings.ToUpper(currency)]
	if !exists {
		return "", fmt.Errorf("currency %s not supported on Solana network %s", currency, network)
	}

	// Empty address means native SOL
	if address == "" {
		return "SOL_NATIVE", nil
	}

	return address, nil
}

// GetDefaultCurrencyForSolanaNetwork 获取Solana网络的默认币种
func GetDefaultCurrencyForSolanaNetwork(cfg *config.Config, network string) string {
	netCfg := GetSolanaNetworkConfig(cfg, network)
	if netCfg != nil {
		return strings.ToUpper(netCfg.NativeToken.Symbol)
	}
	return "SOL"
}
