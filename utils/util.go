package utils

import (
	"fmt"
	"strings"
	"tinypay-server/config"
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

// GetEVMTokenMappingByNetwork 获取EVM代币地址映射 (支持多网络)
func GetEVMTokenMappingByNetwork(cfg *config.Config, network string) map[string]string {
	switch strings.ToLower(network) {
	case "eth-sepolia":
		return map[string]string{
			"ETH":  "0x0000000000000000000000000000000000000000", // Native ETH
			"USDC": strings.ToLower(cfg.ETHSepoliaUSDCAddress),   // Test USDC
		}
	case "celo-sepolia":
		return GetCeloTokenMapping(cfg)
	default:
		// 默认返回以太坊Sepolia映射以保持向后兼容性
		return map[string]string{
			"ETH":  "0x0000000000000000000000000000000000000000", // Native ETH
			"USDC": strings.ToLower(cfg.ETHSepoliaUSDCAddress),   // Test USDC
		}
	}
}

// GetEVMTokenMapping 获取EVM代币地址映射 (向后兼容，默认以太坊Sepolia)
func GetEVMTokenMapping(cfg *config.Config) map[string]string {
	return GetEVMTokenMappingByNetwork(cfg, "eth-sepolia")
}

// GetEVMTokenAddressByNetwork 根据币种获取EVM代币地址 (支持多网络)
func GetEVMTokenAddressByNetwork(cfg *config.Config, currency string, network string) (string, error) {
	if currency == "" {
		// 根据网络设置默认币种
		switch strings.ToLower(network) {
		case "celo-sepolia":
			currency = "CELO"
		default:
			currency = "ETH"
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
	return GetEVMTokenAddressByNetwork(cfg, currency, "eth-sepolia")
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
	return GetCurrencyFromEVMTokenAddressByNetwork(cfg, tokenAddress, "eth-sepolia")
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

// GetCeloTokenMapping 获取Celo Sepolia代币地址映射
func GetCeloTokenMapping(cfg *config.Config) map[string]string {
	return map[string]string{
		"CELO": "0x0000000000000000000000000000000000000000", // Native CELO token
		"USDC": strings.ToLower(cfg.CeloSepoliaUSDCAddress),  // USDC token on Celo Sepolia
	}
}

// GetCeloTokenAddress 根据币种获取Celo Sepolia代币地址
func GetCeloTokenAddress(cfg *config.Config, currency string) (string, error) {
	if currency == "" {
		currency = "CELO" // 默认为CELO
	}

	tokenMapping := GetCeloTokenMapping(cfg)
	tokenAddress, exists := tokenMapping[strings.ToUpper(currency)]
	if !exists {
		return "", fmt.Errorf("unsupported Celo currency: %s", currency)
	}

	return strings.ToLower(tokenAddress), nil
}

// GetCurrencyFromCeloTokenAddress 根据Celo Sepolia代币地址获取币种名称
func GetCurrencyFromCeloTokenAddress(cfg *config.Config, tokenAddress string) string {
	target := strings.ToLower(strings.TrimSpace(tokenAddress))
	tokenMapping := GetCeloTokenMapping(cfg)
	for currency, addr := range tokenMapping {
		if strings.ToLower(strings.TrimSpace(addr)) == target {
			return currency
		}
	}
	return "UNKNOWN"
}

// NetworkCurrencyValidationMatrix 定义网络-货币组合验证矩阵
type NetworkCurrencyValidationMatrix struct {
	validCombinations map[string][]string
}

// NewNetworkCurrencyValidationMatrix 创建新的网络-货币验证矩阵
func NewNetworkCurrencyValidationMatrix() *NetworkCurrencyValidationMatrix {
	return &NetworkCurrencyValidationMatrix{
		validCombinations: map[string][]string{
			"aptos-testnet": {"APT", "USDC"},
			"eth-sepolia":   {"ETH", "USDC"},
			"celo-sepolia":  {"CELO", "USDC"},
		},
	}
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
func ValidateNetworkCurrencyCombination(network, currency string) error {
	matrix := NewNetworkCurrencyValidationMatrix()

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
func GetDefaultCurrencyForNetwork(network string) string {
	switch strings.ToLower(network) {
	case "aptos-testnet":
		return "APT"
	case "eth-sepolia":
		return "ETH"
	case "celo-sepolia":
		return "CELO"
	default:
		return ""
	}
}

// IsNativeCurrency 检查货币是否为指定网络的原生货币
func IsNativeCurrency(network, currency string) bool {
	nativeCurrencies := map[string]string{
		"aptos-testnet": "APT",
		"eth-sepolia":   "ETH",
		"celo-sepolia":  "CELO",
	}

	if nativeCurrency, exists := nativeCurrencies[strings.ToLower(network)]; exists {
		return strings.EqualFold(currency, nativeCurrency)
	}
	return false
}

// ValidateTokenAddressForNetworkCurrency 验证代币地址是否与网络-货币组合匹配
func ValidateTokenAddressForNetworkCurrency(cfg *config.Config, network, currency, tokenAddress string) error {
	// 首先验证网络-货币组合是否有效
	if err := ValidateNetworkCurrencyCombination(network, currency); err != nil {
		return err
	}

	// 获取预期的代币地址
	var expectedAddress string
	var err error

	switch strings.ToLower(network) {
	case "aptos-testnet":
		// Aptos 使用 metadata 地址而不是代币地址
		expectedAddress, err = GetMetadataAddress(cfg, currency)
	case "eth-sepolia", "celo-sepolia":
		expectedAddress, err = GetEVMTokenAddressByNetwork(cfg, currency, network)
	default:
		return fmt.Errorf("unsupported network: %s", network)
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
