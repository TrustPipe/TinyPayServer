package handlers

import (
	"encoding/hex"
	"net/http"
	"strconv"

	"tinypay-server/client"
	"tinypay-server/utils"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	aptosClient *client.AptosClient
}

func NewHandler(aptosClient *client.AptosClient) *Handler {
	return &Handler{
		aptosClient: aptosClient,
	}
}

// PrecommitRequest represents the request body for merchant precommit
type PrecommitRequest struct {
	CommitHash string `json:"commit_hash" binding:"required"` // Hex string
}

// PrecommitResponse represents the response for merchant precommit
type PrecommitResponse struct {
	Success         bool   `json:"success"`
	TransactionHash string `json:"transaction_hash,omitempty"`
	Message         string `json:"message,omitempty"`
	MerchantAddress string `json:"merchant_address"`
}

// CompletePaymentRequest represents the request body for payment completion
type CompletePaymentRequest struct {
	Opt        string `json:"opt" binding:"required"`         // Hex string of opt value
	Payer      string `json:"payer" binding:"required"`       // Payer address
	Recipient  string `json:"recipient" binding:"required"`   // Recipient address
	Amount     string `json:"amount" binding:"required"`      // Amount as string to handle large numbers
	CommitHash string `json:"commit_hash" binding:"required"` // Hex string of commit hash
}

// CompletePaymentResponse represents the response for payment completion
type CompletePaymentResponse struct {
	Success         bool   `json:"success"`
	TransactionHash string `json:"transaction_hash,omitempty"`
	Message         string `json:"message,omitempty"`
	CallerAddress   string `json:"caller_address"`
}

// ComputeHashRequest represents the request body for hash computation
type ComputeHashRequest struct {
	Payer     string `json:"payer" binding:"required"`
	Recipient string `json:"recipient" binding:"required"`
	Amount    string `json:"amount" binding:"required"`
	Opt       string `json:"opt" binding:"required"`
}

// ComputeHashResponse represents the response for hash computation
type ComputeHashResponse struct {
	Success bool   `json:"success"`
	Hash    string `json:"hash,omitempty"`
	Message string `json:"message,omitempty"`
}

// PaymentRequest represents the request body for creating a payment
type PaymentRequest struct {
	PayerAddr string `json:"payer_addr" binding:"required"` // 付款地址 hex格式
	Opt       string `json:"opt" binding:"required"`        // OPT hex格式
	PayeeAddr string `json:"payee_addr" binding:"required"` // 收款地址 hex格式
	Amount    uint64 `json:"amount" binding:"required"`     // 金额 uint类型
	Currency  string `json:"currency"`                      // 货币种类
}

// PaymentResponse represents the response for payment creation
type PaymentResponse struct {
	Status          string   `json:"status,omitempty"`
	TransactionHash string   `json:"transaction_hash,omitempty"`
	Message         string   `json:"message,omitempty"`
	Error           string   `json:"error,omitempty"`
	MissingFields   []string `json:"missing_fields,omitempty"`
	Details         string   `json:"details,omitempty"`
}

// TransactionStatusResponse represents the response for transaction status query
type TransactionStatusResponse struct {
	Status          string `json:"status"`
	TransactionHash string `json:"transaction_hash"`
	Success         *bool  `json:"success,omitempty"`
	ReceivedAmount  string `json:"received_amount,omitempty"`
	Currency        string `json:"currency,omitempty"`
	Message         string `json:"message"`
	Error           string `json:"error,omitempty"`
}

// MerchantPrecommit handles the merchant precommit request
func (h *Handler) MerchantPrecommit(c *gin.Context) {
	var req PrecommitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, PrecommitResponse{
			Success: false,
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	// Convert hex string to bytes
	commitHashBytes, err := hex.DecodeString(req.CommitHash)
	if err != nil {
		c.JSON(http.StatusBadRequest, PrecommitResponse{
			Success: false,
			Message: "Invalid commit hash format: " + err.Error(),
		})
		return
	}

	// Execute merchant precommit
	txHash, err := h.aptosClient.MerchantPrecommit(commitHashBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, PrecommitResponse{
			Success:         false,
			Message:         "Failed to execute precommit: " + err.Error(),
			MerchantAddress: h.aptosClient.GetMerchantAddress(),
		})
		return
	}

	c.JSON(http.StatusOK, PrecommitResponse{
		Success:         true,
		TransactionHash: txHash,
		Message:         "Precommit successful",
		MerchantAddress: h.aptosClient.GetMerchantAddress(),
	})
}

// CompletePayment handles the payment completion request
func (h *Handler) CompletePayment(c *gin.Context) {
	var req CompletePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, CompletePaymentResponse{
			Success: false,
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	// Convert hex strings to bytes
	optBytes := utils.HexToASCIIBytes(req.Opt)

	commitHashBytes, err := hex.DecodeString(req.CommitHash)
	if err != nil {
		c.JSON(http.StatusBadRequest, CompletePaymentResponse{
			Success: false,
			Message: "Invalid commit hash format: " + err.Error(),
		})
		return
	}

	// Convert amount string to uint64
	amount, err := strconv.ParseUint(req.Amount, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, CompletePaymentResponse{
			Success: false,
			Message: "Invalid amount format: " + err.Error(),
		})
		return
	}

	// Execute payment completion
	txHash, err := h.aptosClient.CompletePayment(optBytes, req.Payer, req.Recipient, amount, commitHashBytes)
	if err != nil {
		callerAddr := h.aptosClient.GetPaymasterAddress()
		if callerAddr == "" {
			callerAddr = h.aptosClient.GetMerchantAddress()
		}

		c.JSON(http.StatusInternalServerError, CompletePaymentResponse{
			Success:       false,
			Message:       "Failed to complete payment: " + err.Error(),
			CallerAddress: callerAddr,
		})
		return
	}

	callerAddr := h.aptosClient.GetPaymasterAddress()
	if callerAddr == "" {
		callerAddr = h.aptosClient.GetMerchantAddress()
	}

	c.JSON(http.StatusOK, CompletePaymentResponse{
		Success:         true,
		TransactionHash: txHash,
		Message:         "Payment completed successfully",
		CallerAddress:   callerAddr,
	})
}

// ComputePaymentHash computes the hash for payment parameters
func (h *Handler) ComputePaymentHash(c *gin.Context) {
	var req ComputeHashRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ComputeHashResponse{
			Success: false,
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	// Convert amount string to uint64
	amount, err := strconv.ParseUint(req.Amount, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ComputeHashResponse{
			Success: false,
			Message: "Invalid amount format: " + err.Error(),
		})
		return
	}

	// Convert opt hex string to bytes
	optBytes, err := hex.DecodeString(req.Opt)
	if err != nil {
		c.JSON(http.StatusBadRequest, ComputeHashResponse{
			Success: false,
			Message: "Invalid opt format: " + err.Error(),
		})
		return
	}

	// Compute hash
	hash, err := h.aptosClient.ComputePaymentHash(req.Payer, req.Recipient, amount, optBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ComputeHashResponse{
			Success: false,
			Message: "Failed to compute hash: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ComputeHashResponse{
		Success: true,
		Hash:    hex.EncodeToString(hash),
		Message: "Hash computed successfully",
	})
}

// CreatePayment handles the payment creation request
func (h *Handler) CreatePayment(c *gin.Context) {
	var req PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// Check for missing fields
		missingFields := []string{}
		if req.PayerAddr == "" {
			missingFields = append(missingFields, "payer_addr")
		}
		if req.Opt == "" {
			missingFields = append(missingFields, "opt")
		}
		if req.PayeeAddr == "" {
			missingFields = append(missingFields, "payee_addr")
		}
		if req.Amount == 0 {
			missingFields = append(missingFields, "amount")
		}

		if len(missingFields) > 0 {
			c.JSON(http.StatusBadRequest, PaymentResponse{
				Error:         "missing_fields",
				Message:       "缺少必需字段",
				MissingFields: missingFields,
			})
			return
		}

		c.JSON(http.StatusBadRequest, PaymentResponse{
			Error:   "invalid_request",
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	// Set default currency to APT if not provided
	if req.Currency == "" {
		req.Currency = "APT"
	}

	// Convert hex strings to bytes
	optBytes, err := hex.DecodeString(req.Opt)
	if err != nil {
		c.JSON(http.StatusBadRequest, PaymentResponse{
			Error:   "simulation_failed",
			Message: "交易不合法，模拟失败",
			Details: "Invalid opt format: " + err.Error(),
		})
		return
	}

	// Simulate the transaction first
	_, _, err = h.aptosClient.SimulatePayment([]byte(req.Opt), req.PayerAddr, req.PayeeAddr, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, PaymentResponse{
			Error:   "simulation_failed",
			Message: "交易不合法，模拟失败",
			Details: err.Error(),
		})
		return
	}

	// Submit the transaction
	txHash, err := h.aptosClient.SubmitPayment(optBytes, req.PayerAddr, req.PayeeAddr, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, PaymentResponse{
			Error:   "simulation_failed",
			Message: "交易不合法，模拟失败",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, PaymentResponse{
		Status:          "submitted",
		TransactionHash: txHash,
		Message:         "交易模拟成功，已提交到区块链",
	})
}

// GetTransactionStatus handles the transaction status query
func (h *Handler) GetTransactionStatus(c *gin.Context) {
	txHash := c.Param("transaction_hash")
	if txHash == "" {
		c.JSON(http.StatusBadRequest, TransactionStatusResponse{
			Status:          "error",
			TransactionHash: "",
			Message:         "Transaction hash is required",
		})
		return
	}

	// Check transaction status
	confirmed, err := h.aptosClient.GetTransactionStatus(txHash)
	if err != nil {
		c.JSON(http.StatusNotFound, TransactionStatusResponse{
			Status:          "not_found",
			TransactionHash: txHash,
			Message:         "交易不存在",
			Error:           "not_found",
		})
		return
	}

	if confirmed {
		success := true
		c.JSON(http.StatusOK, TransactionStatusResponse{
			Status:          "confirmed",
			TransactionHash: txHash,
			Success:         &success,
			ReceivedAmount:  "0", // This would need to be extracted from transaction details
			Currency:        "APT",
			Message:         "交易已经被区块链确认",
		})
	} else {
		c.JSON(http.StatusOK, TransactionStatusResponse{
			Status:          "pending",
			TransactionHash: txHash,
			Message:         "交易正在处理中",
		})
	}
}

// HealthCheck provides a simple health check endpoint
func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":            "healthy",
		"merchant_address":  h.aptosClient.GetMerchantAddress(),
		"paymaster_address": h.aptosClient.GetPaymasterAddress(),
	})
}
