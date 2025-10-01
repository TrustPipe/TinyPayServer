package client

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"

	tinypaybindings "tinypay-server/binds/tinypay"
	"tinypay-server/config"
	"tinypay-server/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EVMClient provides helpers to interact with the TinyPay Solidity contract.
type EVMClient struct {
	cfg        *config.Config
	ethClient  *ethclient.Client
	contract   *tinypaybindings.Tinypay
	privateKey *ecdsa.PrivateKey
	chainID    *big.Int
	from       common.Address
	network    string // Track which network this client is configured for
}

// EVMNetworkConfig holds network-specific configuration parameters
type EVMNetworkConfig struct {
	RPCURL          string
	ChainID         uint64
	ContractAddress string
	PrivateKey      string
	Network         string
	NativeToken     config.EVMNativeToken
	Tokens          []config.EVMToken
}

// getNetworkConfig extracts network-specific configuration based on network type
func getNetworkConfig(cfg *config.Config, network string) (*EVMNetworkConfig, error) {
	if cfg == nil {
		return nil, errors.New("config is required")
	}

	// First try to find network in the new EVMNetworks array
	for _, evmNetwork := range cfg.EVMNetworks {
		if evmNetwork.Name == network {
			return &EVMNetworkConfig{
				RPCURL:          evmNetwork.RPCURL,
				ChainID:         evmNetwork.ChainID,
				ContractAddress: evmNetwork.ContractAddress,
				PrivateKey:      evmNetwork.PrivateKey,
				Network:         network,
				NativeToken:     evmNetwork.NativeToken,
				Tokens:          evmNetwork.Tokens,
			}, nil
		}
	}

	return nil, fmt.Errorf("unsupported EVM network: %s", network)
}

// NewEVMClient creates a new EVM client using configuration defaults.
// Deprecated: Use NewEVMClientForNetwork instead for network-specific configuration.
func NewEVMClient(cfg *config.Config) (*EVMClient, error) {
	// Use the first configured EVM network if available, otherwise error
	if len(cfg.EVMNetworks) == 0 {
		return nil, errors.New("no EVM networks configured")
	}
	return NewEVMClientForNetwork(cfg, cfg.EVMNetworks[0].Name)
}

// NewCeloSepoliaClient creates a new EVM client specifically configured for Celo Sepolia.
// Deprecated: use NewEVMClientForNetwork with a configured network name
func NewCeloSepoliaClient(cfg *config.Config) (*EVMClient, error) { return nil, errors.New("deprecated: use NewEVMClientForNetwork") }

// ValidateCeloSepoliaConfig validates Celo Sepolia specific configuration parameters.
// Deprecated: network-specific validation should use generic checks in NewEVMClientForNetwork
func ValidateCeloSepoliaConfig(cfg *config.Config) error { return nil }

// IsCeloSepoliaConfigured checks if Celo Sepolia is properly configured.
// Deprecated
func IsCeloSepoliaConfigured(cfg *config.Config) bool { return false }

// TryNewCeloSepoliaClient attempts to create a Celo Sepolia client with graceful error handling.
// Returns nil and logs warnings if configuration is incomplete, rather than failing.
// Deprecated
func TryNewCeloSepoliaClient(cfg *config.Config) (*EVMClient, error) { return nil, errors.New("deprecated") }

// NewEVMClientForNetwork creates a new EVM client for a specific network.
func NewEVMClientForNetwork(cfg *config.Config, network string) (*EVMClient, error) {

	// Get network-specific configuration
	netCfg, err := getNetworkConfig(cfg, network)
	if err != nil {
		return nil, err
	}

	// Validate network-specific configuration
	if strings.TrimSpace(netCfg.RPCURL) == "" {
		return nil, fmt.Errorf("%s RPC URL is required", strings.ToUpper(strings.Replace(network, "-", "_", -1)))
	}
	if strings.TrimSpace(netCfg.ContractAddress) == "" {
		return nil, fmt.Errorf("%s contract address is required", strings.ToUpper(strings.Replace(network, "-", "_", -1)))
	}
	if strings.TrimSpace(netCfg.PrivateKey) == "" {
		return nil, fmt.Errorf("%s private key is required", strings.ToUpper(strings.Replace(network, "-", "_", -1)))
	}
	if netCfg.ChainID == 0 {
		return nil, fmt.Errorf("%s chain ID must be greater than 0", strings.ToUpper(strings.Replace(network, "-", "_", -1)))
	}

	client, err := ethclient.Dial(netCfg.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s RPC: %w", network, err)
	}

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(netCfg.PrivateKey, "0x"))
	if err != nil {
		return nil, fmt.Errorf("invalid %s private key: %w", network, err)
	}

	publicKey := privateKey.Public()
	ecdsaPubKey, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to parse %s public key", network)
	}

	fromAddress := crypto.PubkeyToAddress(*ecdsaPubKey)

	chainID := big.NewInt(int64(netCfg.ChainID))

	contractAddress := common.HexToAddress(ensureHexPrefix(netCfg.ContractAddress))
	contract, err := tinypaybindings.NewTinypay(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind Tinypay contract for %s: %w", network, err)
	}

	return &EVMClient{
		cfg:        cfg,
		ethClient:  client,
		contract:   contract,
		privateKey: privateKey,
		chainID:    chainID,
		from:       fromAddress,
		network:    network,
	}, nil
}

// Close releases underlying network resources.
func (c *EVMClient) Close() error {
	if c.ethClient != nil {
		c.ethClient.Close()
	}
	return nil
}

// GetConfig returns the configuration used by this client
func (c *EVMClient) GetConfig() *config.Config {
	return c.cfg
}

// GetNetwork returns the network this client is configured for
func (c *EVMClient) GetNetwork() string {
	return c.network
}

// CompletePayment executes the TinyPay completePayment function on the EVM contract.
// optString will be converted to contract bytes using UTF-8 encoding, equivalent to
// ethers.hexlify(ethers.toUtf8Bytes(optString)) in the TypeScript example.
func (c *EVMClient) CompletePayment(
	ctx context.Context,
	tokenAddress string,
	payerAddress string,
	recipientAddress string,
	amount *big.Int,
	optString string,
	commitHashHex string,
) (common.Hash, error) {
	if c == nil {
		return common.Hash{}, errors.New("EVMClient is nil")
	}
	if amount == nil {
		return common.Hash{}, errors.New("amount cannot be nil")
	}

	token := common.HexToAddress(ensureHexPrefix(tokenAddress))
	payer := common.HexToAddress(ensureHexPrefix(payerAddress))
	recipient := common.HexToAddress(ensureHexPrefix(recipientAddress))

	tailBytes := []byte(optString)

	commitHashBytes, err := hexutil.Decode(ensureHexPrefix(commitHashHex))
	if err != nil {
		return common.Hash{}, fmt.Errorf("invalid commit hash: %w", err)
	}
	if len(commitHashBytes) != 32 {
		return common.Hash{}, fmt.Errorf("commit hash must be 32 bytes, got %d", len(commitHashBytes))
	}

	var commitHash [32]byte
	copy(commitHash[:], commitHashBytes)

	auth, err := bind.NewKeyedTransactorWithChainID(c.privateKey, c.chainID)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to create transactor: %w", err)
	}
	auth.From = c.from
	auth.Context = ctx

	log.Printf("auth from %s:", auth.From.String())

	tx, err := c.contract.CompletePayment(auth, token, tailBytes, payer, recipient, amount, commitHash)
	if err != nil {
		return common.Hash{}, fmt.Errorf("completePayment call failed: %w", err)
	}

	return tx.Hash(), nil
}

func ensureHexPrefix(value string) string {
	value = strings.TrimSpace(value)
	if strings.HasPrefix(value, "0x") || strings.HasPrefix(value, "0X") {
		return value
	}
	return "0x" + value
}

// GetTransactionDetails retrieves EVM transaction status and extracts amount from PaymentCompleted event when available.
func (c *EVMClient) GetTransactionDetails(ctx context.Context, txHashHex string) (*TransactionInfo, error) {
	if c == nil || c.ethClient == nil || c.contract == nil {
		return nil, errors.New("EVM client not initialized")
	}
	// Normalize hash
	hash := common.HexToHash(ensureHexPrefix(txHashHex))

	// Check pending status first
	_, isPending, err := c.ethClient.TransactionByHash(ctx, hash)
	if err != nil {
		// Not found or RPC error
		return nil, fmt.Errorf("transaction not found: %w", err)
	}
	if isPending {
		return &TransactionInfo{Confirmed: false}, nil
	}

	// Fetch receipt
	receipt, err := c.ethClient.TransactionReceipt(ctx, hash)
	if err != nil {
		return nil, fmt.Errorf("failed to get receipt: %w", err)
	}

	info := &TransactionInfo{Confirmed: true, Success: receipt.Status == 1}

	// Try to parse PaymentCompleted event to get amount and token
	for _, lg := range receipt.Logs {
		if lg == nil {
			continue
		}
		// Attempt to parse; ignore errors for non-matching topics
		if evt, err := c.contract.ParsePaymentCompleted(*lg); err == nil {
			// Amount
			if evt.Amount != nil {
				info.Amount = evt.Amount.Uint64()
			}
			// CoinType: zero address means native token
			if (evt.Token == common.Address{}) {
				// Set native token symbol based on dynamic network config
				if netCfg := utils.GetEVMNetworkConfig(c.cfg, c.network); netCfg != nil {
					info.CoinType = netCfg.NativeToken.Symbol
				} else {
					info.CoinType = "UNKNOWN"
				}
				info.TokenAddress = ""
			} else {
				info.TokenAddress = evt.Token.Hex()
				// Map token address to currency using network-specific utility function
				if currency := utils.GetCurrencyFromEVMTokenAddressByNetwork(c.cfg, evt.Token.Hex(), c.network); currency != "" {
					info.CoinType = currency
				} else {
					// Fallback to native token based on dynamic network config
					if netCfg := utils.GetEVMNetworkConfig(c.cfg, c.network); netCfg != nil {
						info.CoinType = netCfg.NativeToken.Symbol
					} else {
						info.CoinType = "UNKNOWN"
					}
				}
			}
			break
		}
	}

	if !info.Success {
		info.Error = "execution failed"
	}
	return info, nil
}

// GetUserLimits returns TinyPay user limits on EVM
func (c *EVMClient) GetUserLimits(ctx context.Context, userAddress string) (*UserLimits, error) {
	if c == nil || c.contract == nil {
		return nil, errors.New("EVM client not initialized")
	}
	addr := common.HexToAddress(ensureHexPrefix(userAddress))
	res, err := c.contract.GetUserLimits(&bind.CallOpts{Context: ctx}, addr)
	if err != nil {
		return nil, fmt.Errorf("getUserLimits failed: %w", err)
	}
	return &UserLimits{
		PaymentLimit:    res.PaymentLimit,
		TailUpdateCount: res.TailUpdateCount,
		MaxTailUpdates:  res.MaxTailUpdates,
	}, nil
}
