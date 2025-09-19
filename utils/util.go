package utils

import (
	"fmt"
	"strings"
)

// HexToASCIIBytes 将十六进制字符串转换为 ASCII 字节数组
func HexToASCIIBytes(hexStr string) []byte {
	// 清理输入
	hexStr = strings.ToLower(strings.TrimSpace(hexStr))

	return []byte(strings.TrimPrefix(hexStr, "0x"))
}

// CoinTypeMapping 币种类型映射
var CoinTypeMapping = map[string]string{
	"APT":  "0x1::aptos_coin::AptosCoin",
	"USDC": "0xaadbf0681ef3dc9decd123340db16954f85319853533ed4ace6ec5d11aaad190::test_usdc::TestUSDC",
}

// GetCoinType 根据币种名称获取对应的合约类型
func GetCoinType(currency string) (string, error) {
	if currency == "" {
		currency = "APT" // 默认为APT
	}
	
	coinType, exists := CoinTypeMapping[strings.ToUpper(currency)]
	if !exists {
		return "", fmt.Errorf("unsupported currency: %s", currency)
	}
	
	return coinType, nil
}

// GetSupportedCurrencies 获取支持的币种列表
func GetSupportedCurrencies() []string {
	currencies := make([]string, 0, len(CoinTypeMapping))
	for currency := range CoinTypeMapping {
		currencies = append(currencies, currency)
	}
	return currencies
}
