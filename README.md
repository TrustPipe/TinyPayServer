# TinyPayServer

一个基于 Aptos 区块链的支付服务器，提供安全、快速的加密货币支付解决方案。

## 功能特性

- 🔐 安全的支付处理
- 🚀 基于 Aptos 区块链
- 📡 RESTful API 接口
- 🔄 实时交易状态查询
- 📚 完整的 API 文档
- 🐳 Docker 容器化部署
- 🌐 Nginx 反向代理和 CORS 支持

## 快速开始

### 环境要求

- Docker 和 Docker Compose
- Go 1.22+ (本地开发)

### 使用 Docker Compose 部署

1. **克隆项目**
   ```bash
   git clone <repository-url>
   cd tinypay-server
   ```

2. **配置环境变量**
   ```bash
   cp .env.example .env
   # 编辑 .env 文件，填入必要的配置
   ```

3. **启动服务**
   ```bash
   # 构建并启动所有服务
   docker-compose up --build

   # 后台运行
   docker-compose up -d --build
   ```

4. **访问服务**
   - API 服务: https://api-tinypay.predictplay.xyz (生产环境) 或 http://localhost (本地开发)
   - API 文档: https://api-tinypay.predictplay.xyz/docs
   - 健康检查: https://api-tinypay.predictplay.xyz/api/health

### 本地开发

1. **安装依赖**
   ```bash
   go mod download
   ```

2. **配置环境变量**
   ```bash
   cp .env.example .env
   # 编辑 .env 文件
   ```

3. **运行服务**
   ```bash
   go run main.go
   ```

## 服务架构

```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│   Client    │───▶│    Nginx    │───▶│ TinyPay     │
│             │    │  (Port 80)  │    │ Server      │
│             │    │             │    │ (Port 9090) │
└─────────────┘    └─────────────┘    └─────────────┘
                           │
                           ▼
                   ┌─────────────┐
                   │   Aptos     │
                   │ Blockchain  │
                   └─────────────┘
```

## 配置说明

### 环境变量

#### 基础配置

| 变量名 | 描述 | 默认值 |
|--------|------|--------|
| `PORT` | 服务端口 | `9090` |

#### Aptos 网络配置

| 变量名 | 描述 | 默认值 |
|--------|------|--------|
| `APTOS_NETWORK` | Aptos 网络 | `testnet` |
| `APTOS_NODE_URL` | Aptos 节点 URL | `https://fullnode.testnet.aptoslabs.com/v1` |
| `APTOS_FAUCET_URL` | Aptos 水龙头 URL | `https://faucet.testnet.aptoslabs.com` |
| `CONTRACT_ADDRESS` | TinyPay 合约地址 | 必填 |
| `USDC_METADATA_ADDRESS` | USDC 元数据地址 | 必填 |
| `MERCHANT_PRIVATE_KEY` | 商户私钥 | 必填 |
| `PAYMASTER_PRIVATE_KEY` | 付费主私钥 | 可选 |

#### Ethereum Sepolia 网络配置

| 变量名 | 描述 | 默认值 |
|--------|------|--------|
| `ETH_SEPOLIA_RPC_URL` | Ethereum Sepolia RPC URL | 必填 |
| `ETH_SEPOLIA_CHAIN_ID` | Ethereum Sepolia 链 ID | `11155111` |
| `ETH_SEPOLIA_CONTRACT_ADDRESS` | TinyPay 合约地址 | 必填 |
| `ETH_SEPOLIA_PRIVATE_KEY` | 私钥 | 必填 |
| `ETH_SEPOLIA_USDC_ADDRESS` | USDC 代币地址 | 必填 |

#### Celo Sepolia 网络配置

| 变量名 | 描述 | 默认值 |
|--------|------|--------|
| `CELO_SEPOLIA_RPC_URL` | Celo Sepolia RPC URL | `https://alfajores-forno.celo-testnet.org` |
| `CELO_SEPOLIA_CHAIN_ID` | Celo Sepolia 链 ID | `44787` |
| `CELO_SEPOLIA_CONTRACT_ADDRESS` | TinyPay 合约地址 | 必填 |
| `CELO_SEPOLIA_PRIVATE_KEY` | 私钥 | 必填 |
| `CELO_SEPOLIA_USDC_ADDRESS` | USDC 代币地址 | `0x2F25deB3848C207fc8E0c34035B3Ba7fC157602B` |

#### 可选配置

| 变量名 | 描述 | 默认值 |
|--------|------|--------|
| `MAX_GAS_AMOUNT` | 最大 Gas 数量 | `2000` |
| `GAS_UNIT_PRICE` | Gas 单价 | `100` |

### Docker 服务

- **tinypay-server**: 主要的支付服务
- **nginx**: 反向代理服务器，提供 CORS 支持和负载均衡

## API 文档

### 主要端点

- `GET /api/health` - 健康检查
- `POST /api/payments` - 创建支付
- `GET /api/payments/{hash}` - 查询交易状态
- `GET /docs` - Swagger UI 文档
- `GET /openapi.yaml` - OpenAPI 规范

### 示例请求

```bash
# 健康检查
curl https://api-tinypay.predictplay.xyz/api/health

# 创建支付
curl -X POST https://api-tinypay.predictplay.xyz/api/payments \
  -H "Content-Type: application/json" \
  -d '{
    "payer_addr": "0x1234...",
    "payee_addr": "0x5678...",
    "amount": 1000000,
    "otp": "deadbeef"
  }'
```

## 开发指南

### 项目结构

```
.
├── api/                 # OpenAPI 生成的代码和规范
├── client/              # Aptos 客户端
├── config/              # 配置管理
├── handlers/            # HTTP 处理器
├── examples/            # 使用示例
├── docker-compose.yml   # Docker Compose 配置
├── Dockerfile          # Docker 镜像构建
├── nginx.conf          # Nginx 配置
└── main.go             # 主程序入口
```

### 构建和测试

```bash
# 构建
go build -o tinypay-server .

# 测试
go test ./...

# 生成 API 代码
make generate
```

## 部署说明

### SSL 证书配置

项目已配置支持 HTTPS，使用 Let's Encrypt 证书：

1. **获取 SSL 证书**
   ```bash
   # 使用 certbot 获取证书
   sudo certbot certonly --nginx -d api-tinypay.predictplay.xyz
   ```

2. **证书路径**
   - 证书文件: `/etc/letsencrypt/live/api-tinypay.predictplay.xyz/fullchain.pem`
   - 私钥文件: `/etc/letsencrypt/live/api-tinypay.predictplay.xyz/privkey.pem`

3. **自动续期**
   ```bash
   # 添加到 crontab
   0 12 * * * /usr/bin/certbot renew --quiet && docker-compose restart nginx
   ```

### 生产环境部署

1. **确保 SSL 证书存在**
   - 证书文件必须存在于指定路径
   - Docker 容器会挂载主机的证书目录

2. **安全配置**
   - 使用 Docker secrets 管理敏感信息
   - 配置防火墙规则
   - 启用日志监控

3. **性能优化**
   - 调整 Nginx 工作进程数
   - 配置连接池和缓存
   - 监控资源使用情况

### 监控和日志

- 应用日志: `./logs/`
- Nginx 日志: `./logs/nginx/`
- 健康检查: `https://api-tinypay.predictplay.xyz/api/health`
- SSL 证书状态: 可通过浏览器或 SSL 检查工具验证

## 故障排除

### 常见问题

1. **服务无法启动**
   - 检查环境变量配置
   - 确认端口未被占用
   - 查看 Docker 日志

2. **CORS 错误**
   - 检查 Nginx 配置
   - 确认请求头设置正确

3. **交易失败**
   - 检查 Aptos 网络连接
   - 验证私钥和地址
   - 查看交易日志

### 查看日志

```bash
# 查看所有服务日志
docker-compose logs

# 查看特定服务日志
docker-compose logs tinypay-server
docker-compose logs nginx

# 实时日志
docker-compose logs -f
```

## 贡献指南

1. Fork 项目
2. 创建功能分支
3. 提交更改
4. 推送到分支
5. 创建 Pull Request

## 许可证

[MIT License](LICENSE)

## 支持

如有问题，请创建 Issue 或联系维护团队。