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

// GetEVMTokenMapping 获取EVM代币地址映射
func GetEVMTokenMapping(cfg *config.Config) map[string]string {
	return map[string]string{
		"ETH":  "0x0000000000000000000000000000000000000000", // Native ETH
		"USDC": strings.ToLower(cfg.EVMTestUSDCAddress), // Test USDC
	}
}

// GetEVMTokenAddress 根据币种获取EVM代币地址
func GetEVMTokenAddress(cfg *config.Config, currency string) (string, error) {
	if currency == "" {
		currency = "ETH" // 默认为ETH
	}

	tokenMapping := GetEVMTokenMapping(cfg)
	tokenAddress, exists := tokenMapping[strings.ToUpper(currency)]
	if !exists {
		return "", fmt.Errorf("unsupported EVM currency: %s", currency)
	}

	return strings.ToLower(tokenAddress), nil
}

// GetCurrencyFromEVMTokenAddress 根据EVM代币地址获取币种名称
func GetCurrencyFromEVMTokenAddress(cfg *config.Config, tokenAddress string) string {
	target := strings.ToLower(strings.TrimSpace(tokenAddress))
	tokenMapping := GetEVMTokenMapping(cfg)
	for currency, addr := range tokenMapping {
		if strings.ToLower(strings.TrimSpace(addr)) == target {
			return currency
		}
	}
	return "UNKNOWN"
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
