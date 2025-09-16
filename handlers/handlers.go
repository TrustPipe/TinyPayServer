package handlers

import (
	"encoding/hex"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"tinypay-server/client"
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
	Opt         string `json:"opt" binding:"required"`         // Hex string of opt value
	Payer       string `json:"payer" binding:"required"`       // Payer address
	Recipient   string `json:"recipient" binding:"required"`   // Recipient address
	Amount      string `json:"amount" binding:"required"`      // Amount as string to handle large numbers
	CommitHash  string `json:"commit_hash" binding:"required"` // Hex string of commit hash
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
	optBytes, err := hex.DecodeString(req.Opt)
	if err != nil {
		c.JSON(http.StatusBadRequest, CompletePaymentResponse{
			Success: false,
			Message: "Invalid opt format: " + err.Error(),
		})
		return
	}

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

// HealthCheck provides a simple health check endpoint
func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":           "healthy",
		"merchant_address": h.aptosClient.GetMerchantAddress(),
		"paymaster_address": h.aptosClient.GetPaymasterAddress(),
	})
}
