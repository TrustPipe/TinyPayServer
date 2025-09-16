package client

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/aptos-labs/aptos-go-sdk"
	"github.com/aptos-labs/aptos-go-sdk/bcs"
	"tinypay-server/config"
)

type AptosClient struct {
	client          *aptos.Client
	config          *config.Config
	merchantAccount *aptos.Account
	paymasterAccount *aptos.Account
}

func NewAptosClient(cfg *config.Config) (*AptosClient, error) {
	// Create network config based on environment
	var networkConfig aptos.NetworkConfig
	switch cfg.AptosNetwork {
	case "devnet":
		networkConfig = aptos.DevnetConfig
	case "testnet":
		networkConfig = aptos.TestnetConfig
	case "mainnet":
		networkConfig = aptos.MainnetConfig
	default:
		// Custom network configuration
		networkConfig = aptos.NetworkConfig{
			NodeUrl:   cfg.AptosNodeURL,
			FaucetUrl: cfg.AptosFaucetURL,
		}
	}

	// Create Aptos client
	client, err := aptos.NewClient(networkConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Aptos client: %w", err)
	}

	// Load merchant account from private key
	merchantAccount, err := aptos.NewEd25519Account()
	if err != nil {
		return nil, fmt.Errorf("failed to create merchant account: %w", err)
	}

	// Load paymaster account if provided
	var paymasterAccount *aptos.Account
	if cfg.PaymasterPrivateKey != "" {
		paymasterAccount, err = aptos.NewEd25519Account()
		if err != nil {
			return nil, fmt.Errorf("failed to create paymaster account: %w", err)
		}
	}

	return &AptosClient{
		client:           client,
		config:           cfg,
		merchantAccount:  merchantAccount,
		paymasterAccount: paymasterAccount,
	}, nil
}

// MerchantPrecommit executes the merchant_precommit function
func (ac *AptosClient) MerchantPrecommit(commitHash []byte) (string, error) {
	log.Printf("Executing merchant_precommit with hash: %x", commitHash)

	// Convert commit_hash to BCS bytes for vector<u8>
	commitHashBytes, err := bcs.SerializeBytes(commitHash)
	if err != nil {
		return "", fmt.Errorf("failed to serialize commit hash: %w", err)
	}

	// Build transaction
	rawTxn, err := ac.client.BuildTransaction(
		ac.merchantAccount.AccountAddress(),
		aptos.TransactionPayload{
			Payload: &aptos.EntryFunction{
				Module: aptos.ModuleId{
					Address: *parseAccountAddress(ac.config.ContractAddress),
					Name:    "tinypay",
				},
				Function: "merchant_precommit",
				ArgTypes: []aptos.TypeTag{},
				Args: [][]byte{
					commitHashBytes,
				},
			},
		},
		aptos.MaxGasAmount(ac.config.MaxGasAmount),
		aptos.GasUnitPrice(ac.config.GasUnitPrice),
	)
	if err != nil {
		return "", fmt.Errorf("failed to build transaction: %w", err)
	}

	// Simulate transaction (optional but recommended)
	simulationResult, err := ac.client.SimulateTransaction(rawTxn, ac.merchantAccount)
	if err != nil {
		log.Printf("Warning: failed to simulate transaction: %v", err)
	} else {
		log.Printf("Simulation - Gas used: %d, Gas unit price: %d, Total fee: %d",
			simulationResult[0].GasUsed,
			simulationResult[0].GasUnitPrice,
			simulationResult[0].GasUsed*simulationResult[0].GasUnitPrice)
	}

	// Sign transaction
	signedTxn, err := rawTxn.SignedTransaction(ac.merchantAccount)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Submit transaction
	submitResult, err := ac.client.SubmitTransaction(signedTxn)
	if err != nil {
		return "", fmt.Errorf("failed to submit transaction: %w", err)
	}

	// Wait for transaction completion
	_, err = ac.client.WaitForTransaction(submitResult.Hash)
	if err != nil {
		return "", fmt.Errorf("failed to wait for transaction: %w", err)
	}

	log.Printf("Merchant precommit successful, transaction hash: %s", submitResult.Hash)
	return submitResult.Hash, nil
}

// CompletePayment executes the complete_payment function
func (ac *AptosClient) CompletePayment(opt []byte, payer, recipient string, amount uint64, commitHash []byte) (string, error) {
	log.Printf("Executing complete_payment - Payer: %s, Recipient: %s, Amount: %d", payer, recipient, amount)

	// Parse addresses
	payerAddr := parseAccountAddress(payer)
	recipientAddr := parseAccountAddress(recipient)

	// Serialize parameters
	optBytes, err := bcs.SerializeBytes(opt)
	if err != nil {
		return "", fmt.Errorf("failed to serialize opt: %w", err)
	}

	payerBytes, err := bcs.Serialize(payerAddr)
	if err != nil {
		return "", fmt.Errorf("failed to serialize payer address: %w", err)
	}

	recipientBytes, err := bcs.Serialize(recipientAddr)
	if err != nil {
		return "", fmt.Errorf("failed to serialize recipient address: %w", err)
	}

	amountBytes, err := bcs.SerializeU64(amount)
	if err != nil {
		return "", fmt.Errorf("failed to serialize amount: %w", err)
	}

	commitHashBytes, err := bcs.SerializeBytes(commitHash)
	if err != nil {
		return "", fmt.Errorf("failed to serialize commit hash: %w", err)
	}

	// Choose the caller (merchant or paymaster)
	var caller *aptos.Account
	if ac.paymasterAccount != nil {
		caller = ac.paymasterAccount
		log.Println("Using paymaster account as caller")
	} else {
		caller = ac.merchantAccount
		log.Println("Using merchant account as caller")
	}

	// Build transaction
	rawTxn, err := ac.client.BuildTransaction(
		caller.AccountAddress(),
		aptos.TransactionPayload{
			Payload: &aptos.EntryFunction{
				Module: aptos.ModuleId{
					Address: *parseAccountAddress(ac.config.ContractAddress),
					Name:    "tinypay",
				},
				Function: "complete_payment",
				ArgTypes: []aptos.TypeTag{},
				Args: [][]byte{
					optBytes,
					payerBytes,
					recipientBytes,
					amountBytes,
					commitHashBytes,
				},
			},
		},
		aptos.MaxGasAmount(ac.config.MaxGasAmount),
		aptos.GasUnitPrice(ac.config.GasUnitPrice),
	)
	if err != nil {
		return "", fmt.Errorf("failed to build transaction: %w", err)
	}

	// Simulate transaction (optional but recommended)
	simulationResult, err := ac.client.SimulateTransaction(rawTxn, caller)
	if err != nil {
		log.Printf("Warning: failed to simulate transaction: %v", err)
	} else {
		log.Printf("Simulation - Gas used: %d, Gas unit price: %d, Total fee: %d",
			simulationResult[0].GasUsed,
			simulationResult[0].GasUnitPrice,
			simulationResult[0].GasUsed*simulationResult[0].GasUnitPrice)
	}

	// Sign transaction
	signedTxn, err := rawTxn.SignedTransaction(caller)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Submit transaction
	submitResult, err := ac.client.SubmitTransaction(signedTxn)
	if err != nil {
		return "", fmt.Errorf("failed to submit transaction: %w", err)
	}

	// Wait for transaction completion
	_, err = ac.client.WaitForTransaction(submitResult.Hash)
	if err != nil {
		return "", fmt.Errorf("failed to wait for transaction: %w", err)
	}

	log.Printf("Payment completion successful, transaction hash: %s", submitResult.Hash)
	return submitResult.Hash, nil
}

// Helper function to compute payment parameters hash
func (ac *AptosClient) ComputePaymentHash(payer, recipient string, amount uint64, opt []byte) ([]byte, error) {
	// Parse addresses
	payerAddr := parseAccountAddress(payer)
	recipientAddr := parseAccountAddress(recipient)

	// Serialize parameters in the same order as the contract
	payerBytes, err := bcs.Serialize(payerAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize payer address: %w", err)
	}

	recipientBytes, err := bcs.Serialize(recipientAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize recipient address: %w", err)
	}

	amountBytes, err := bcs.SerializeU64(amount)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize amount: %w", err)
	}

	optBytes, err := bcs.SerializeBytes(opt)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize opt: %w", err)
	}

	// Concatenate all parameters
	var paramsBytes []byte
	paramsBytes = append(paramsBytes, payerBytes...)
	paramsBytes = append(paramsBytes, recipientBytes...)
	paramsBytes = append(paramsBytes, amountBytes...)
	paramsBytes = append(paramsBytes, optBytes...)

	// Compute SHA256 hash
	hash := sha256.Sum256(paramsBytes)
	return hash[:], nil
}

// Helper function to convert bytes to hex string
func (ac *AptosClient) BytesToHex(data []byte) string {
	return hex.EncodeToString(data)
}

// Helper function to convert hex string to bytes
func (ac *AptosClient) HexToBytes(hexStr string) ([]byte, error) {
	return hex.DecodeString(hexStr)
}

// GetMerchantAddress returns the merchant account address
func (ac *AptosClient) GetMerchantAddress() string {
	return ac.merchantAccount.Address.String()
}

// GetPaymasterAddress returns the paymaster account address (if available)
func (ac *AptosClient) GetPaymasterAddress() string {
	if ac.paymasterAccount != nil {
		return ac.paymasterAccount.Address.String()
	}
	return ""
}

// Helper function to parse account address
func parseAccountAddress(addr string) *aptos.AccountAddress {
	address := &aptos.AccountAddress{}
	err := address.ParseStringRelaxed(addr)
	if err != nil {
		panic("Failed to parse address: " + addr + ", error: " + err.Error())
	}
	return address
}
