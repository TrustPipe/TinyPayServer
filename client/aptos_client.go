package client

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"strings"

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

// SimulatePayment simulates a payment transaction without submitting it
func (ac *AptosClient) SimulatePayment(opt []byte, payer, recipient string, amount uint64) error {
	log.Printf("Simulating payment - Payer: %s, Recipient: %s, Amount: %d", payer, recipient, amount)

	// Parse addresses
	payerAddr := parseAccountAddress(payer)
	recipientAddr := parseAccountAddress(recipient)

	// Serialize parameters
	optBytes, err := bcs.SerializeBytes(opt)
	if err != nil {
		return fmt.Errorf("failed to serialize opt: %w", err)
	}

	payerBytes, err := bcs.Serialize(payerAddr)
	if err != nil {
		return fmt.Errorf("failed to serialize payer address: %w", err)
	}

	recipientBytes, err := bcs.Serialize(recipientAddr)
	if err != nil {
		return fmt.Errorf("failed to serialize recipient address: %w", err)
	}

	amountBytes, err := bcs.SerializeU64(amount)
	if err != nil {
		return fmt.Errorf("failed to serialize amount: %w", err)
	}

	// Compute commit hash for simulation
	commitHash, err := ac.ComputePaymentHash(payer, recipient, amount, opt)
	if err != nil {
		return fmt.Errorf("failed to compute commit hash: %w", err)
	}

	commitHashBytes, err := bcs.SerializeBytes(commitHash)
	if err != nil {
		return fmt.Errorf("failed to serialize commit hash: %w", err)
	}

	// Choose the caller (merchant or paymaster)
	var caller *aptos.Account
	if ac.paymasterAccount != nil {
		caller = ac.paymasterAccount
	} else {
		caller = ac.merchantAccount
	}

	// Build transaction for simulation
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
		return fmt.Errorf("failed to build transaction: %w", err)
	}

	// Simulate transaction
	simulationResult, err := ac.client.SimulateTransaction(rawTxn, caller)
	if err != nil {
		return fmt.Errorf("transaction simulation failed: %w", err)
	}

	// Check if simulation was successful
	if len(simulationResult) == 0 || !simulationResult[0].Success {
		if len(simulationResult) > 0 {
			return fmt.Errorf("transaction simulation failed: %s", simulationResult[0].VmStatus)
		}
		return fmt.Errorf("transaction simulation failed: unknown error")
	}

	log.Printf("Simulation successful - Gas used: %d, Gas unit price: %d, Total fee: %d",
		simulationResult[0].GasUsed,
		simulationResult[0].GasUnitPrice,
		simulationResult[0].GasUsed*simulationResult[0].GasUnitPrice)

	return nil
}

// TransactionInfo contains detailed transaction information
type TransactionInfo struct {
	Confirmed bool
	Success   bool
	Amount    uint64
	Error     string
}

// GetTransactionStatus gets the status of a transaction by hash
func (ac *AptosClient) GetTransactionStatus(txHash string) (bool, error) {
	log.Printf("Getting transaction status for hash: %s", txHash)

	// Try to wait for transaction to check if it exists and is confirmed
	_, err := ac.client.WaitForTransaction(txHash)
	if err != nil {
		// Transaction might be pending or failed
		return false, nil
	}

	// Transaction is confirmed
	return true, nil
}

// GetTransactionDetails gets detailed information about a transaction
func (ac *AptosClient) GetTransactionDetails(txHash string) (*TransactionInfo, error) {
	log.Printf("Getting transaction details for hash: %s", txHash)

	// Try to wait for transaction to check if it exists and is confirmed
	txnResult, err := ac.client.WaitForTransaction(txHash)
	if err != nil {
		// Transaction might be pending or failed
		return &TransactionInfo{
			Confirmed: false,
			Success:   false,
			Amount:    0,
			Error:     err.Error(),
		}, nil
	}

	// Transaction is confirmed, extract amount from events
	amount := uint64(0)
	success := txnResult.Success

	// Parse transaction events to extract the actual amount
	log.Printf("Transaction has %d events", len(txnResult.Events))
	if txnResult.Events != nil {
		for i, event := range txnResult.Events {
			log.Printf("Event %d: Type=%s, Data=%+v", i, event.Type, event.Data)
			
			// Prioritize fungible_asset events for amount extraction
			if event.Type == "0x1::fungible_asset::Withdraw" || event.Type == "0x1::fungible_asset::Deposit" {
				if amountStr, exists := event.Data["amount"]; exists {
					log.Printf("Found amount in fungible_asset event: %v (type: %T)", amountStr, amountStr)
					if amountFloat, ok := amountStr.(float64); ok {
						amount = uint64(amountFloat)
						log.Printf("Parsed fungible_asset amount as float64: %d", amount)
						break // Use the first fungible_asset amount found
					} else if amountString, ok := amountStr.(string); ok {
						if parsedAmount, parseErr := strconv.ParseUint(amountString, 10, 64); parseErr == nil {
							amount = parsedAmount
							log.Printf("Parsed fungible_asset amount as string: %d", amount)
							break // Use the first fungible_asset amount found
						}
					}
				}
			}
		}
		
		// If no fungible_asset amount found, try other event types
		if amount == 0 {
			for i, event := range txnResult.Events {
				log.Printf("Fallback check Event %d: Type=%s", i, event.Type)
				
				// Look for other coin transfer events
				if event.Type == "0x1::coin::WithdrawEvent" || 
				   event.Type == "0x1::coin::DepositEvent" ||
				   event.Type == "0x1::aptos_coin::WithdrawEvent" ||
				   event.Type == "0x1::aptos_coin::DepositEvent" ||
				   strings.Contains(event.Type, "PaymentCompleted") {
					
					// Try to extract amount from event data
					if amountStr, exists := event.Data["amount"]; exists {
						log.Printf("Found amount in fallback event: %v (type: %T)", amountStr, amountStr)
						if amountFloat, ok := amountStr.(float64); ok {
							amount = uint64(amountFloat)
							log.Printf("Parsed fallback amount as float64: %d", amount)
							break
						} else if amountString, ok := amountStr.(string); ok {
							if parsedAmount, parseErr := strconv.ParseUint(amountString, 10, 64); parseErr == nil {
								amount = parsedAmount
								log.Printf("Parsed fallback amount as string: %d", amount)
								break
							}
						}
					}
				}
			}
		}
	}
	
	log.Printf("Final extracted amount: %d octas", amount)

	return &TransactionInfo{
		Confirmed: true,
		Success:   success,
		Amount:    amount,
		Error:     "",
	}, nil
}

// SubmitPayment creates and submits a payment transaction
func (ac *AptosClient) SubmitPayment(opt []byte, payer, recipient string, amount uint64) (string, error) {
	log.Printf("Submitting payment - Payer: %s, Recipient: %s, Amount: %d", payer, recipient, amount)

	// First simulate the transaction
	err := ac.SimulatePayment(opt, payer, recipient, amount)
	if err != nil {
		return "", fmt.Errorf("simulation failed: %w", err)
	}

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

	// Compute commit hash
	commitHash, err := ac.ComputePaymentHash(payer, recipient, amount, opt)
	if err != nil {
		return "", fmt.Errorf("failed to compute commit hash: %w", err)
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

	log.Printf("Payment submitted successfully, transaction hash: %s", submitResult.Hash)
	return submitResult.Hash, nil
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
