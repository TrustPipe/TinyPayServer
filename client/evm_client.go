package client

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strings"

	tinypaybindings "tinypay-server/binds/tinypay"
	"tinypay-server/config"

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
}

// NewEVMClient creates a new EVM client using configuration defaults.
func NewEVMClient(cfg *config.Config) (*EVMClient, error) {
	if cfg == nil {
		return nil, errors.New("config is required")
	}
	if strings.TrimSpace(cfg.EVMRPCURL) == "" {
		return nil, errors.New("EVM_RPC_URL is required")
	}
	if strings.TrimSpace(cfg.EVMContractAddress) == "" {
		return nil, errors.New("EVM_CONTRACT_ADDRESS is required")
	}
	if strings.TrimSpace(cfg.EVMPrivateKey) == "" {
		return nil, errors.New("EVM_PRIVATE_KEY is required")
	}
	if cfg.EVMChainID == 0 {
		return nil, errors.New("EVM_CHAIN_ID must be greater than 0")
	}

	client, err := ethclient.Dial(cfg.EVMRPCURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to EVM RPC: %w", err)
	}

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(cfg.EVMPrivateKey, "0x"))
	if err != nil {
		return nil, fmt.Errorf("invalid EVM private key: %w", err)
	}

	publicKey := privateKey.Public()
	ecdsaPubKey, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("failed to parse EVM public key")
	}

	fromAddress := crypto.PubkeyToAddress(*ecdsaPubKey)

	chainID := big.NewInt(int64(cfg.EVMChainID))

	contractAddress := common.HexToAddress(ensureHexPrefix(cfg.EVMContractAddress))
	contract, err := tinypaybindings.NewTinypay(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind Tinypay contract: %w", err)
	}

	return &EVMClient{
		cfg:        cfg,
		ethClient:  client,
		contract:   contract,
		privateKey: privateKey,
		chainID:    chainID,
		from:       fromAddress,
	}, nil
}

// Close releases underlying network resources.
func (c *EVMClient) Close() error {
	if c.ethClient != nil {
		c.ethClient.Close()
	}
	return nil
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
