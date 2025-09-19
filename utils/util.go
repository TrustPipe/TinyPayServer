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

// GetCoinTypeMapping 根据配置获取币种类型映射
func GetCoinTypeMapping(cfg *config.Config) map[string]string {
	return map[string]string{
		"APT":  "0x1::aptos_coin::AptosCoin",
		"USDC": cfg.USDCContractAddress,
	}
}

// GetCoinType 根据币种名称获取对应的合约类型
func GetCoinType(cfg *config.Config, currency string) (string, error) {
	if currency == "" {
		currency = "APT" // 默认为APT
	}
	
	coinTypeMapping := GetCoinTypeMapping(cfg)
	coinType, exists := coinTypeMapping[strings.ToUpper(currency)]
	if !exists {
		return "", fmt.Errorf("unsupported currency: %s", currency)
	}
	
	return coinType, nil
}

// GetSupportedCurrencies 获取支持的币种列表
func GetSupportedCurrencies(cfg *config.Config) []string {
	coinTypeMapping := GetCoinTypeMapping(cfg)
	currencies := make([]string, 0, len(coinTypeMapping))
	for currency := range coinTypeMapping {
		currencies = append(currencies, currency)
	}
	return currencies
}

// GetCurrencyFromCoinType 根据合约类型获取币种名称（反向映射）
func GetCurrencyFromCoinType(cfg *config.Config, coinType string) string {
	coinTypeMapping := GetCoinTypeMapping(cfg)
	for currency, contractType := range coinTypeMapping {
		if contractType == coinType {
			return currency
		}
	}
	return "UNKNOWN"
}
