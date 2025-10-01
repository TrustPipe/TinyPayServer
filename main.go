package main

import (
	"log"

	"tinypay-server/api"
	"tinypay-server/client"
	"tinypay-server/config"

	"github.com/gin-gonic/gin"
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

	// Initialize EVM clients for different networks
	evmClients := make(map[string]*client.EVMClient)

	// Initialize EVM clients from the new EVMNetworks configuration
	for _, evmNetwork := range cfg.EVMNetworks {
		evmClient, err := client.NewEVMClientForNetwork(cfg, evmNetwork.Name)
		if err != nil {
			log.Printf("Warning: Failed to initialize %s client: %v. %s payments will not be available.", evmNetwork.Name, err, evmNetwork.Name)
		} else {
			evmClients[evmNetwork.Name] = evmClient
			log.Printf("%s client initialized successfully", evmNetwork.Name)
		}
	}

	log.Printf("Merchant address: %s", aptosClient.GetMerchantAddress())
	if paymasterAddr := aptosClient.GetPaymasterAddress(); paymasterAddr != "" {
		log.Printf("Paymaster address: %s", paymasterAddr)
	}

	// Initialize OpenAPI server
	apiServer := api.NewAPIServer(aptosClient, evmClients, cfg)

	// Setup Gin router
	router := gin.Default()

	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
		}

		c.Next()
	})

	// OpenAPI generated routes
	api.RegisterHandlers(router, apiServer)

	// Setup API documentation
	api.SetupDocumentationRoutes(router)

	// Start server
	log.Printf("Server starting on :%s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
