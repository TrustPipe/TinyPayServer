package utils

import "strings"

// HexToASCIIBytes 将十六进制字符串转换为 ASCII 字节数组
func HexToASCIIBytes(hexStr string) []byte {
	// 清理输入
	hexStr = strings.ToLower(strings.TrimSpace(hexStr))

	return []byte(strings.TrimPrefix(hexStr, "0x"))
}
