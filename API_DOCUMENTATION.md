# Payment API 接口文档

## 接口概述

基于 RESTful 规范的支付接口，支持多区块链网络（Aptos、Ethereum）交易处理。

## 接口列表

### 1. 创建支付交易

**POST** `/api/payments`

创建新的支付交易，支持多网络（Aptos、Ethereum）。服务器会先检查字段完整性，然后进行交易模拟，模拟成功后返回交易哈希并上链提交。

#### 请求参数

```json
{
  "payer_addr": "string",  // 付款地址 hex格式
  "otp": "string",         // OPT hex格式 (仅 Aptos 网络需要)
  "payee_addr": "string",  // 收款地址 hex格式
  "amount": number,        // 金额 uint类型
  "currency": "string",    // 货币种类 (APT/ETH/CELO/USDC)
  "network": "string"      // 目标网络 (aptos-testnet/eth-sepolia/celo-sepolia，默认 aptos-testnet)
}
```

#### 响应参数

**模拟成功，交易已提交 (200)**
```json
{
  "code": 1001,
  "data": {
    "transaction_hash": "string", // 交易哈希
    "network": "string"           // 网络标识
  }
}
```

**字段缺失 (400)**
```json
{
  "code": 2004,
  "data": null
}
```

**交易模拟失败 (400)**
```json
{
  "code": 2002,
  "data": null
}
```

**无效的货币种类 (400)**
```json
{
  "code": 2006,
  "data": null
}
```

**无效的网络参数 (400)**
```json
{
  "code": 2003,
  "data": null
}
```

---

### 2. 查询交易状态

**GET** `/api/payments/{transaction_hash}?network={network}`

根据交易哈希查询交易状态和详情，支持多网络查询。

#### 路径参数

- `transaction_hash`: 交易哈希值

#### 查询参数

- `network`: 目标网络 (可选，默认 `aptos-testnet`)
  - `aptos-testnet`: Aptos 测试网
  - `eth-sepolia`: Ethereum Sepolia 测试网
  - `celo-sepolia`: Celo Sepolia 测试网

#### 响应参数

**交易处理中 (200)**
```json
{
  "code": 1002,
  "data": {
    "status": "pending",
    "network": "aptos-testnet"
  }
}
```

**交易确认成功 (200)**
```json
{
  "code": 1003,
  "data": {
    "status": "confirmed",
    "received_amount": 1000000,  // 实际收到的金额 (整数)
    "currency": "APT",           // 货币种类 (APT/ETH)
    "network": "aptos-testnet"   // 网络标识
  }
}
```

**交易确认失败 (200)**
```json
{
  "code": 1003,
  "data": {
    "status": "failed",
    "error": "string",           // 失败原因
    "network": "aptos-testnet"
  }
}
```

**交易不存在 (404)**
```json
{
  "code": 2005,
  "data": null
}
```

**无效的网络参数 (400)**
```json
{
  "code": 2003,
  "data": null
}
```

---

### 3. 查询用户限制

**GET** `/api/users/{user_address}/limits?network={network}`

根据用户地址查询用户的支付限制信息，支持多网络查询。

#### 路径参数

- `user_address`: 用户地址 hex格式

#### 查询参数

- `network`: 目标网络 (可选，默认 `aptos-testnet`)
  - `aptos-testnet`: Aptos 测试网
  - `eth-sepolia`: Ethereum Sepolia 测试网
  - `celo-sepolia`: Celo Sepolia 测试网

#### 响应参数

**查询成功 (200)**
```json
{
  "code": 1000,
  "data": {
    "user_limits": {
      "payment_limit": 1000000,     // 支付限额
      "tail_update_count": 5,       // 尾部更新次数
      "max_tail_updates": 10        // 最大尾部更新次数
    },
    "network": "aptos-testnet"      // 网络标识
  }
}
```

**用户地址无效 (400)**
```json
{
  "code": 2003,
  "data": null
}
```

**无效的网络参数 (400)**
```json
{
  "code": 2003,
  "data": null
}
```

## 使用流程

### 支付流程

1. 前端调用 `POST /api/payments` 创建支付交易，指定目标网络
2. 服务器检查字段完整性和网络参数有效性
3. 服务器根据网络类型进行交易模拟验证
4. 模拟成功则返回交易哈希并提交到区块链，模拟失败则返回错误信息
5. 前端收到交易哈希后，使用 `GET /api/payments/{transaction_hash}?network={network}` 轮询查询状态
6. 等待状态变为 `confirmed` 或 `failed`

### 用户限制查询流程

1. 前端调用 `GET /api/users/{user_address}/limits?network={network}` 查询用户限制信息
2. 服务器验证用户地址格式和网络参数
3. 根据网络类型调用对应区块链合约的view函数获取用户限制数据
4. 返回包含支付限额、尾部更新次数等信息的响应

## 网络支持

### Aptos 网络 (aptos-testnet)
- 支持货币：APT, USDC
- 需要 OPT 参数进行交易验证
- 交易哈希格式：0x + 64位十六进制

### Ethereum 网络 (eth-sepolia)
- 支持货币：ETH, USDC
- 无需 OPT 参数
- 交易哈希格式：0x + 64位十六进制

### Celo 网络 (celo-sepolia)
- 支持货币：CELO, USDC
- 无需 OPT 参数
- 交易哈希格式：0x + 64位十六进制
- 网络特性：EVM 兼容，支持原生 CELO 代币和 USDC ERC-20 代币

## 示例请求

### 创建 Aptos 支付
```bash
curl -X POST "http://localhost:9090/api/payments" \
  -H "Content-Type: application/json" \
  -d '{
    "payer_addr": "0x1234...",
    "otp": "0xabcd...",
    "payee_addr": "0x5678...",
    "amount": 1000000,
    "currency": "APT",
    "network": "aptos-testnet"
  }'
```

### 创建 Ethereum 支付
```bash
curl -X POST "http://localhost:9090/api/payments" \
  -H "Content-Type: application/json" \
  -d '{
    "payer_addr": "0x1234...",
    "payee_addr": "0x5678...",
    "amount": 1000000,
    "currency": "ETH",
    "network": "eth-sepolia"
  }'
```

### 创建 Celo 支付 (CELO 原生代币)
```bash
curl -X POST "http://localhost:9090/api/payments" \
  -H "Content-Type: application/json" \
  -d '{
    "payer_addr": "0x1234...",
    "payee_addr": "0x5678...",
    "amount": 1000000,
    "currency": "CELO",
    "network": "celo-sepolia"
  }'
```

### 创建 Celo 支付 (USDC 代币)
```bash
curl -X POST "http://localhost:9090/api/payments" \
  -H "Content-Type: application/json" \
  -d '{
    "payer_addr": "0x1234...",
    "payee_addr": "0x5678...",
    "amount": 1000000,
    "currency": "USDC",
    "network": "celo-sepolia"
  }'
```

### 查询交易状态
```bash
# Aptos 网络
curl "http://localhost:9090/api/payments/0x1234...?network=aptos-testnet"

# Ethereum 网络
curl "http://localhost:9090/api/payments/0x1234...?network=eth-sepolia"

# Celo 网络
curl "http://localhost:9090/api/payments/0x1234...?network=celo-sepolia"
```

### 查询用户限制
```bash
# Aptos 网络
curl "http://localhost:9090/api/users/0x1234.../limits?network=aptos-testnet"

# Ethereum 网络
curl "http://localhost:9090/api/users/0x1234.../limits?network=eth-sepolia"

# Celo 网络
curl "http://localhost:9090/api/users/0x1234.../limits?network=celo-sepolia"
```