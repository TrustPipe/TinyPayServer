package client

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"tinypay-server/config"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

// SolanaClient provides methods to interact with TinyPay Solana program
type SolanaClient struct {
	client    *rpc.Client
	config    *config.Config
	programID solana.PublicKey
	paymaster solana.PrivateKey
	network   string
}

// NewSolanaClient creates a new Solana client for the specified network
func NewSolanaClient(cfg *config.Config, network string) (*SolanaClient, error) {
	// Get network-specific configuration
	var netCfg *config.SolanaNetwork
	for i := range cfg.SolanaNetworks {
		if cfg.SolanaNetworks[i].Name == network {
			netCfg = &cfg.SolanaNetworks[i]
			break
		}
	}

	if netCfg == nil {
		return nil, fmt.Errorf("solana network %s not found in configuration", network)
	}

	// Validate configuration
	if netCfg.RPCURL == "" {
		return nil, fmt.Errorf("solana RPC URL is required for network %s", network)
	}
	if netCfg.ProgramID == "" {
		return nil, fmt.Errorf("solana program ID is required for network %s", network)
	}
	if netCfg.PaymasterPrivateKey == "" {
		return nil, fmt.Errorf("solana paymaster private key is required for network %s", network)
	}

	// Create RPC client
	client := rpc.New(netCfg.RPCURL)

	// Parse program ID
	programID, err := solana.PublicKeyFromBase58(netCfg.ProgramID)
	if err != nil {
		return nil, fmt.Errorf("invalid program ID: %w", err)
	}

	// Parse paymaster private key
	paymasterKey, err := solana.PrivateKeyFromBase58(netCfg.PaymasterPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("invalid paymaster private key: %w", err)
	}

	return &SolanaClient{
		client:    client,
		config:    cfg,
		programID: programID,
		paymaster: paymasterKey,
		network:   network,
	}, nil
}

// GetConfig returns the configuration
func (sc *SolanaClient) GetConfig() *config.Config {
	return sc.config
}

// GetNetwork returns the network name
func (sc *SolanaClient) GetNetwork() string {
	return sc.network
}

// GetPaymasterAddress returns the paymaster public key as a string
func (sc *SolanaClient) GetPaymasterAddress() string {
	return sc.paymaster.PublicKey().String()
}

// ConvertOTPForContract converts OTP string to bytes for contract
func ConvertOTPForContract(otpString string) []byte {
	return []byte(otpString)
}

// ComputeOTPHash computes SHA256 hash of OTP and returns hex ASCII bytes
func ComputeOTPHash(otp []byte) []byte {
	hash := sha256.Sum256(otp)
	hexString := hex.EncodeToString(hash[:])
	return []byte(hexString)
}

// CompletePayment executes the complete_payment instruction on Solana
func (sc *SolanaClient) CompletePayment(
	ctx context.Context,
	payerPubkey solana.PublicKey,
	recipientPubkey solana.PublicKey,
	otpString string,
	amountLamports uint64,
) (solana.Signature, error) {
	log.Printf("Executing Solana complete_payment - Payer: %s, Recipient: %s, Amount: %d", 
		payerPubkey.String(), recipientPubkey.String(), amountLamports)

	// Convert OTP to bytes
	otpBytes := ConvertOTPForContract(otpString)

	// Derive PDAs
	userAccountPDA, _, err := solana.FindProgramAddress(
		[][]byte{
			[]byte("user"),
			payerPubkey.Bytes(),
		},
		sc.programID,
	)
	if err != nil {
		return solana.Signature{}, fmt.Errorf("failed to derive user account PDA: %w", err)
	}

	statePDA, _, err := solana.FindProgramAddress(
		[][]byte{[]byte("state")},
		sc.programID,
	)
	if err != nil {
		return solana.Signature{}, fmt.Errorf("failed to derive state PDA: %w", err)
	}

	vaultPDA, _, err := solana.FindProgramAddress(
		[][]byte{[]byte("vault")},
		sc.programID,
	)
	if err != nil {
		return solana.Signature{}, fmt.Errorf("failed to derive vault PDA: %w", err)
	}

	// Build instruction data
	instructionData := buildCompletePaymentInstruction(otpBytes, amountLamports)

	// Build transaction instruction
	instruction := solana.NewInstruction(
		sc.programID,
		solana.AccountMetaSlice{
			solana.Meta(userAccountPDA).WRITE(),
			solana.Meta(statePDA).WRITE(),
			solana.Meta(vaultPDA).WRITE(),
			solana.Meta(recipientPubkey).WRITE(),
			solana.Meta(sc.paymaster.PublicKey()).SIGNER(),
			solana.Meta(solana.SystemProgramID),
		},
		instructionData,
	)

	// Get latest blockhash (replaces deprecated GetRecentBlockhash)
	recent, err := sc.client.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		return solana.Signature{}, fmt.Errorf("failed to get latest blockhash: %w", err)
	}

	// Build and sign transaction
	tx, err := solana.NewTransaction(
		[]solana.Instruction{instruction},
		recent.Value.Blockhash,
		solana.TransactionPayer(sc.paymaster.PublicKey()),
	)
	if err != nil {
		return solana.Signature{}, fmt.Errorf("failed to create transaction: %w", err)
	}

	_, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		if key.Equals(sc.paymaster.PublicKey()) {
			return &sc.paymaster
		}
		return nil
	})
	if err != nil {
		return solana.Signature{}, fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Send transaction
	sig, err := sc.client.SendTransaction(ctx, tx)
	if err != nil {
		return solana.Signature{}, fmt.Errorf("failed to send transaction: %w", err)
	}

	log.Printf("Solana payment completed! Signature: %s", sig)
	return sig, nil
}

// computeAnchorDiscriminator computes the Anchor discriminator for an instruction
// Discriminator = first 8 bytes of SHA256("global:function_name")
func computeAnchorDiscriminator(functionName string) []byte {
	preimage := "global:" + functionName
	hash := sha256.Sum256([]byte(preimage))
	return hash[:8]
}

// buildCompletePaymentInstruction constructs the instruction data for complete_payment
// Format: [discriminator(8)] + [otp_len(4)] + [otp] + [amount(8)]
func buildCompletePaymentInstruction(otp []byte, amount uint64) []byte {
	// Calculate the correct discriminator for "complete_payment"
	discriminator := computeAnchorDiscriminator("complete_payment")

	data := make([]byte, 0)
	data = append(data, discriminator...)

	// OTP length (4 bytes, little-endian)
	otpLen := make([]byte, 4)
	binary.LittleEndian.PutUint32(otpLen, uint32(len(otp)))
	data = append(data, otpLen...)

	// OTP data
	data = append(data, otp...)

	// Amount (8 bytes, little-endian)
	amountBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(amountBytes, amount)
	data = append(data, amountBytes...)

	return data
}

// GetUserLimits queries user limits from the Solana program
func (sc *SolanaClient) GetUserLimits(ctx context.Context, userPubkey solana.PublicKey) (*UserLimits, error) {
	log.Printf("Getting Solana user limits for address: %s", userPubkey.String())

	// Derive user account PDA
	userAccountPDA, _, err := solana.FindProgramAddress(
		[][]byte{
			[]byte("user"),
			userPubkey.Bytes(),
		},
		sc.programID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to derive user account PDA: %w", err)
	}

	// Get account info
	accountInfo, err := sc.client.GetAccountInfo(ctx, userAccountPDA)
	if err != nil {
		return nil, fmt.Errorf("failed to get account info: %w", err)
	}

	if accountInfo == nil || accountInfo.Value == nil {
		return nil, fmt.Errorf("user account not found")
	}

	// Deserialize user account
	userAccount, err := deserializeSolanaUserAccount(accountInfo.Value.Data.GetBinary())
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize user account: %w", err)
	}

	return &UserLimits{
		PaymentLimit:    userAccount.PaymentLimit,
		TailUpdateCount: userAccount.TailUpdateCount,
		MaxTailUpdates:  userAccount.MaxTailUpdates,
	}, nil
}

// SolanaUserAccount represents the user account structure in Solana program
type SolanaUserAccount struct {
	Owner           solana.PublicKey
	Balance         uint64
	Tail            [64]byte
	PaymentLimit    uint64
	TailUpdateCount uint64
	MaxTailUpdates  uint64
	Bump            uint8
}

// deserializeSolanaUserAccount deserializes user account data
// Layout: [8 bytes discriminator] + [account data]
func deserializeSolanaUserAccount(data []byte) (*SolanaUserAccount, error) {
	if len(data) < 8+137 { // 8 (discriminator) + 137 (account data)
		return nil, fmt.Errorf("invalid account data length: %d", len(data))
	}

	// Skip 8-byte discriminator
	offset := 8

	account := &SolanaUserAccount{}

	// Owner (32 bytes)
	copy(account.Owner[:], data[offset:offset+32])
	offset += 32

	// Balance (8 bytes, little-endian)
	account.Balance = binary.LittleEndian.Uint64(data[offset : offset+8])
	offset += 8

	// Tail (64 bytes)
	copy(account.Tail[:], data[offset:offset+64])
	offset += 64

	// PaymentLimit (8 bytes, little-endian)
	account.PaymentLimit = binary.LittleEndian.Uint64(data[offset : offset+8])
	offset += 8

	// TailUpdateCount (8 bytes, little-endian)
	account.TailUpdateCount = binary.LittleEndian.Uint64(data[offset : offset+8])
	offset += 8

	// MaxTailUpdates (8 bytes, little-endian)
	account.MaxTailUpdates = binary.LittleEndian.Uint64(data[offset : offset+8])
	offset += 8

	// Bump (1 byte)
	account.Bump = data[offset]

	return account, nil
}

// GetTransactionDetails retrieves Solana transaction status and details
func (sc *SolanaClient) GetTransactionDetails(ctx context.Context, signature string) (*TransactionInfo, error) {
	log.Printf("Getting Solana transaction details for signature: %s", signature)

	// Parse signature
	sig, err := solana.SignatureFromBase58(signature)
	if err != nil {
		return nil, fmt.Errorf("invalid signature format: %w", err)
	}

	// Get transaction with details
	maxSupportedTransactionVersion := uint64(0)
	out, err := sc.client.GetTransaction(
		ctx,
		sig,
		&rpc.GetTransactionOpts{
			Encoding:                       solana.EncodingBase64,
			MaxSupportedTransactionVersion: &maxSupportedTransactionVersion,
			Commitment:                     rpc.CommitmentConfirmed,
		},
	)

	if err != nil {
		return nil, fmt.Errorf("transaction not found: %w", err)
	}

	if out == nil {
		return &TransactionInfo{
			Confirmed: false,
		}, nil
	}

	// Check if transaction succeeded
	success := out.Meta.Err == nil
	
	var errorMsg string
	if out.Meta.Err != nil {
		errorMsg = fmt.Sprintf("%v", out.Meta.Err)
	}

	// Extract amount from transaction by parsing instruction data
	amount := uint64(0)
	if out.Transaction != nil {
		tx, err := out.Transaction.GetTransaction()
		if err == nil && len(tx.Message.Instructions) > 0 {
			// Get the first instruction (complete_payment)
			instruction := tx.Message.Instructions[0]
			if len(instruction.Data) >= 20 { // discriminator(8) + otp_len(4) + amount(8)
				// Amount is at the end of instruction data (last 8 bytes)
				amountOffset := len(instruction.Data) - 8
				amount = binary.LittleEndian.Uint64(instruction.Data[amountOffset:])
			}
		}
	}

	return &TransactionInfo{
		Confirmed: true,
		Success:   success,
		Amount:    amount,
		CoinType:  "SOL", // Default to SOL, can be extended for SPL tokens
		Error:     errorMsg,
	}, nil
}

// Close releases resources (if any)
func (sc *SolanaClient) Close() error {
	// No resources to release for now
	return nil
}
