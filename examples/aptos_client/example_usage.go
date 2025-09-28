package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"tinypay-server/client"
	"tinypay-server/config"
)

// 这个示例展示了如何直接使用 AptosClient 进行合约调用
func main() {
	// 加载配置
	cfg := config.LoadConfig()

	// 初始化 Aptos 客户端
	aptosClient, err := client.NewAptosClient(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize Aptos client: %v", err)
	}

	fmt.Printf("Merchant Address: %s\n", aptosClient.GetMerchantAddress())
	if paymasterAddr := aptosClient.GetPaymasterAddress(); paymasterAddr != "" {
		fmt.Printf("Paymaster Address: %s\n", paymasterAddr)
	}

	// 示例 1: 计算支付参数哈希
	fmt.Println("\n=== 示例 1: 计算支付参数哈希 ===")
	payer := "0x1234567890abcdef1234567890abcdef12345678"
	recipient := "0xabcdef1234567890abcdef1234567890abcdef12"
	amount := uint64(1000000) // 1 APT (假设精度为 8)
	otp := "previous_iteration_hex_string"

	optBytes, _ := hex.DecodeString(otp)
	hash, err := aptosClient.ComputePaymentHash(payer, recipient, amount, optBytes)
	if err != nil {
		log.Fatalf("Failed to compute hash: %v", err)
	}

	fmt.Printf("Payment Hash: %s\n", hex.EncodeToString(hash))

	// 示例 2: 商户预提交 (需要实际的网络连接)
	if os.Getenv("RUN_NETWORK_EXAMPLES") == "true" {
		fmt.Println("\n=== 示例 2: 商户预提交 ===")
		txHash, err := aptosClient.MerchantPrecommit(hash)
		if err != nil {
			log.Printf("Precommit failed: %v", err)
		} else {
			fmt.Printf("Precommit Transaction Hash: %s\n", txHash)

			// 示例 3: 完成支付
			fmt.Println("\n=== 示例 3: 完成支付 ===")
			txHash2, err := aptosClient.CompletePayment(optBytes, payer, recipient, amount, hash)
			if err != nil {
				log.Printf("Payment completion failed: %v", err)
			} else {
				fmt.Printf("Payment Transaction Hash: %s\n", txHash2)
			}
		}
	} else {
		fmt.Println("\n提示: 设置环境变量 RUN_NETWORK_EXAMPLES=true 来运行网络相关的示例")
	}
}
