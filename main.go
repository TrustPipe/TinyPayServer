package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"tinypay-server/client"
	"tinypay-server/config"
	"tinypay-server/handlers"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()
	log.Printf("Starting TinyPay server on port %s", cfg.Port)
	log.Printf("Contract address: %s", cfg.ContractAddress)
	log.Printf("Network: %s", cfg.AptosNetwork)

	// Initialize Aptos client
	aptosClient, err := client.NewAptosClient(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize Aptos client: %v", err)
	}

	log.Printf("Merchant address: %s", aptosClient.GetMerchantAddress())
	if paymasterAddr := aptosClient.GetPaymasterAddress(); paymasterAddr != "" {
		log.Printf("Paymaster address: %s", paymasterAddr)
	}

	// Initialize handlers
	handler := handlers.NewHandler(aptosClient)

	// Setup Gin router
	router := gin.Default()

	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// API routes
	api := router.Group("/api/v1")
	{
		// Health check
		api.GET("/health", handler.HealthCheck)

		// Merchant precommit endpoint
		api.POST("/merchant/precommit", handler.MerchantPrecommit)

		// Payment completion endpoint
		api.POST("/payment/complete", handler.CompletePayment)

		// Utility endpoint to compute payment hash
		api.POST("/utils/compute-hash", handler.ComputePaymentHash)
	}

	// Start server
	log.Printf("Server starting on :%s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
