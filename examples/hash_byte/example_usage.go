package main

import (
	"crypto/sha256"
	"fmt"
	"tinypay-server/utils"
)

// VerifyHashChain 验证哈希链的正确性
func VerifyHashChain(optHex, expectedTailHex string) bool {
	// 将 opt 转换为 ASCII 字节
	optASCIIBytes := utils.HexToASCIIBytes(optHex)

	// 对 opt ASCII 字节进行 SHA256 哈希
	hash := sha256.Sum256(optASCIIBytes)

	// 将哈希结果转换为十六进制字符串
	computedHex := fmt.Sprintf("%x", hash)

	return computedHex == expectedTailHex
}

func main() {
	// 测试数据（来自你的 Python 脚本输出）
	optHex := "84eb882e56142984dea2fee9772d60c05d3885941fd2522761451446f46ae437"
	tailHex := "adb6beedc72be327ccbc58cf8c866ea608603c27568ec0752dc7d1e7608507a6"

	// 转换为 ASCII 字节数组
	optASCIIBytes := utils.HexToASCIIBytes(optHex)
	tailASCIIBytes := utils.HexToASCIIBytes(tailHex)

	fmt.Printf("opt hex: %s\n", optHex)
	fmt.Printf("opt ASCII bytes: %v\n", optASCIIBytes)
	fmt.Printf("opt ASCII bytes length: %d\n", len(optASCIIBytes))
	fmt.Println()

	fmt.Printf("tail hex: %s\n", tailHex)
	fmt.Printf("tail ASCII bytes: %v\n", tailASCIIBytes)
	fmt.Printf("tail ASCII bytes length: %d\n", len(tailASCIIBytes))
	fmt.Println()

	// 验证哈希链
	isValid := VerifyHashChain(optHex, tailHex)
	fmt.Printf("Hash chain verification: %t\n", isValid)

	// 显示用于 Aptos CLI 的格式
	fmt.Printf("\nAptos CLI format for opt:\n")
	fmt.Printf("u8:[")
	for i, b := range optASCIIBytes {
		if i > 0 {
			fmt.Printf(",")
		}
		fmt.Printf("%d", b)
	}
	fmt.Printf("]\n")
}
