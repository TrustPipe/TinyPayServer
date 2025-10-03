# TinyPayServer

A blockchain-based payment server built on Aptos and EVM networks, providing secure and fast cryptocurrency payment solutions with multi-network support.

## Features

- 🔐 Secure payment processing
- 🚀 Multi-blockchain support (Aptos, Ethereum, Celo)
- 📡 RESTful API with OpenAPI 3.0 specification
- 🔄 Real-time transaction status tracking
- 📚 Comprehensive API documentation with Swagger UI
- 🐳 Docker containerization
- 🌐 Nginx reverse proxy with CORS support
- ⚙️ Flexible configuration system (TOML/ENV)
- 🔧 Code generation with oapi-codegen

## Quick Start

### Prerequisites

- Docker and Docker Compose
- Go 1.22+ (for local development)

### Docker Deployment

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd tinypay-server
   ```

2. **Configure environment**
   ```bash
   cp .env.example .env
   # Edit .env file with your configuration
   # OR
   cp config.toml.example config.toml
   # Edit config.toml file (recommended for new deployments)
   ```

3. **Start services**
   ```bash
   # Build and start all services
   docker-compose up --build

   # Run in background
   docker-compose up -d --build
   ```

4. **Access services**
   - API Service: http://localhost:9090
   - API Documentation: http://localhost:9090/docs
   - Health Check: http://localhost:9090/api/health
   - OpenAPI Spec: http://localhost:9090/openapi.yaml

### Local Development

1. **Install dependencies**
   ```bash
   go mod download
   ```

2. **Install development tools**
   ```bash
   make install
   ```

3. **Configure environment**
   ```bash
   cp config.toml.example config.toml
   # Edit config.toml with your settings
   ```

4. **Generate API code**
   ```bash
   make generate
   ```

5. **Run the server**
   ```bash
   make dev
   # OR
   go run .
   ```

## Architecture

```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│   Client    │───▶│    Nginx    │───▶│ TinyPay     │
│             │    │  (Port 80)  │    │ Server      │
│             │    │             │    │ (Port 9090) │
└─────────────┘    └─────────────┘    └─────────────┘
                           │
                           ▼
                   ┌─────────────┐
                   │ Multi-Chain │
                   │ Blockchain  │
                   │ Networks    │
                   └─────────────┘
```

## Configuration

### Configuration Methods

The server supports two configuration methods:

1. **TOML Configuration (Recommended)**: Use `config.toml` for structured configuration
2. **Environment Variables**: Use `.env` file for legacy compatibility

### TOML Configuration

Create a `config.toml` file based on `config.toml.example`:

```toml
# Aptos Network Configuration
[aptos]
network = "testnet"
node_url = "https://fullnode.testnet.aptoslabs.com/v1"
faucet_url = "https://faucet.testnet.aptoslabs.com"

# Contract Configuration
[contract]
address = "0x5877584f4dbd72b5d101f32be3bea1eb67e96020ded3943919ddc80927c88893"
usdc_metadata_address = "0x69091fbab5f7d635ee7ac5098cf0c1efbe31d68fec0f2cd565e8d168daf52832"

# Server Configuration
[server]
port = "9090"

# Gas Configuration
[gas]
max_gas_amount = 100000
gas_unit_price = 100

# Private Keys
[keys]
merchant_private_key = "0x..."
paymaster_private_key = "0x..."

# EVM Networks Configuration
[[evm_networks]]
name = "eth-sepolia"
rpc_url = "https://sepolia.infura.io/v3/YOUR_PROJECT_ID"
chain_id = 11155111
contract_address = "0x..."
private_key = "0x..."

[evm_networks.native_token]
symbol = "ETH"
address = "0x0000000000000000000000000000000000000000"

[[evm_networks.tokens]]
symbol = "USDC"
address = "0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238"

[[evm_networks]]
name = "celo-sepolia"
rpc_url = "https://alfajores-forno.celo-testnet.org"
chain_id = 44787
contract_address = "0x..."
private_key = "0x..."

[evm_networks.native_token]
symbol = "CELO"
address = "0x0000000000000000000000000000000000000000"

[[evm_networks.tokens]]
symbol = "USDC"
address = "0x01C5C0122039549AD1493B8220cABEdD739BC44E"
```

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `9090` |
| `APTOS_NETWORK` | Aptos network | `testnet` |
| `APTOS_NODE_URL` | Aptos node URL | `https://fullnode.testnet.aptoslabs.com/v1` |
| `CONTRACT_ADDRESS` | TinyPay contract address | Required |
| `MERCHANT_PRIVATE_KEY` | Merchant private key | Required |
| `ETH_SEPOLIA_RPC_URL` | Ethereum Sepolia RPC URL | Required |
| `ETH_SEPOLIA_CONTRACT_ADDRESS` | Ethereum contract address | Required |
| `CELO_SEPOLIA_RPC_URL` | Celo Sepolia RPC URL | `https://alfajores-forno.celo-testnet.org` |

## API Documentation

### Supported Networks

- **aptos-testnet**: Aptos testnet
- **eth-sepolia**: Ethereum Sepolia testnet
- **celo-sepolia**: Celo Sepolia testnet

### Supported Currencies

- **APT**: Aptos native token
- **ETH**: Ethereum native token
- **CELO**: Celo native token
- **USDC**: USD Coin (available on all networks)

### Main Endpoints

- `GET /api/health` - Health check
- `POST /api/payments` - Create payment transaction
- `GET /api/payments/{hash}?network={network}` - Query transaction status
- `GET /docs` - Swagger UI documentation
- `GET /openapi.yaml` - OpenAPI specification

### Response Format

All API responses follow a unified format:

```json
{
  "code": 1000,
  "data": { /* response data or null */ }
}
```

### Status Codes

#### Success Codes (1000-1999)
- `1000`: Server running normally
- `1001`: Transaction created successfully
- `1002`: Transaction processing
- `1003`: Transaction confirmed

#### Error Codes (2000-2999)
- `2000`: Amount must be greater than 0
- `2001`: Amount exceeds limit
- `2002`: Insufficient balance
- `2003`: Invalid OTP
- `2004`: Missing required fields
- `2005`: Transaction not found
- `2006`: Invalid currency type

### Example Requests

```bash
# Health check
curl http://localhost:9090/api/health

# Create payment (Aptos)
curl -X POST http://localhost:9090/api/payments \
  -H "Content-Type: application/json" \
  -d '{
    "payer_addr": "0x1234...",
    "payee_addr": "0x5678...",
    "amount": 1000000,
    "currency": "USDC",
    "network": "aptos-testnet",
    "otp": "deadbeef"
  }'

# Create payment (Ethereum)
curl -X POST http://localhost:9090/api/payments \
  -H "Content-Type: application/json" \
  -d '{
    "payer_addr": "0x1234...",
    "payee_addr": "0x5678...",
    "amount": 1000000,
    "currency": "USDC",
    "network": "eth-sepolia"
  }'

# Query transaction status
curl "http://localhost:9090/api/payments/0xabc123...?network=aptos-testnet"
```

## Development

### Project Structure

```
.
├── api/                    # Generated API code and OpenAPI spec
│   ├── openapi.yaml       # OpenAPI 3.0 specification
│   ├── server.gen.go      # Generated server interfaces
│   ├── types.gen.go       # Generated data types
│   ├── client.gen.go      # Generated client code
│   └── spec.gen.go        # Generated spec embedding
├── client/                # Blockchain client implementations
│   ├── aptos_client.go    # Aptos blockchain client
│   └── evm_client.go      # EVM blockchain client
├── config/                # Configuration management
│   ├── config.go          # Configuration loading logic
│   └── config_test.go     # Configuration tests
├── cmd/                   # Command-line tools and utilities
├── examples/              # Usage examples
├── binds/                 # Smart contract bindings
├── utils/                 # Utility functions
├── main.go               # Application entry point
├── Makefile              # Build automation
├── docker-compose.yml    # Docker services configuration
└── Dockerfile           # Container image definition
```

### Build Commands

```bash
# Install development tools
make install

# Generate API code from OpenAPI spec
make generate

# Build the application
make build

# Run in development mode
make dev

# Run tests
make test

# Clean generated files
make clean

# View available commands
make help
```

### Code Generation

The project uses **Design-First API development** with OpenAPI 3.0:

1. Define API in `api/openapi.yaml`
2. Generate Go code with `oapi-codegen`
3. Implement business logic in handlers

### Dependencies

Key dependencies:
- **github.com/getkin/kin-openapi**: OpenAPI 3.0 specification processing
- **github.com/oapi-codegen/runtime**: Runtime utilities for generated code
- **github.com/gin-gonic/gin**: HTTP web framework
- **github.com/pelletier/go-toml/v2**: TOML configuration parsing

## Deployment

### Docker Services

- **tinypay-server**: Main payment service
- **nginx**: Reverse proxy with CORS support

### Production Deployment

1. **SSL Configuration**
   ```bash
   # Obtain SSL certificate
   sudo certbot certonly --nginx -d your-domain.com
   ```

2. **Security Setup**
   - Use Docker secrets for sensitive data
   - Configure firewall rules
   - Enable monitoring and logging

3. **Performance Optimization**
   - Adjust Nginx worker processes
   - Configure connection pooling
   - Monitor resource usage

### Health Monitoring

- Application logs: `./logs/`
- Health endpoint: `/api/health`
- Docker health checks enabled

## Troubleshooting

### Common Issues

1. **Service won't start**
   - Check configuration files
   - Verify port availability
   - Review Docker logs

2. **CORS errors**
   - Check Nginx configuration
   - Verify request headers

3. **Transaction failures**
   - Check blockchain network connectivity
   - Verify private keys and addresses
   - Review transaction logs

### Viewing Logs

```bash
# View all service logs
docker-compose logs

# View specific service logs
docker-compose logs tinypay-server

# Follow logs in real-time
docker-compose logs -f
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

[MIT License](LICENSE)

## Support

For issues and questions, please create an issue in the repository or contact the development team.