# Payment API 接口文档

## 接口概述

基于 RESTful 规范的支付接口，支持 Aptos 区块链交易处理。

## 接口列表

### 1. 创建支付交易

**POST** `/api/payments`

创建新的支付交易，服务器会先检查字段完整性，然后进行交易模拟，模拟成功后返回交易哈希并上链提交。

#### 请求参数

```json
{
  "payer_addr": "string",  // 付款地址 hex格式
  "opt": "string",         // OPT hex格式
  "payee_addr": "string",  // 收款地址 hex格式
  "amount": number,        // 金额 uint类型
  "currency": "string"     // 货币种类
}
```

#### 响应参数

**模拟成功，交易已提交 (200)**
```json
{
  "status": "submitted",
  "transaction_hash": "string", // 交易哈希
  "message": "交易模拟成功，已提交到区块链"
}
```

**字段缺失 (400)**
```json
{
  "error": "missing_fields",
  "message": "缺少必需字段",
  "missing_fields": ["string"] // 缺失的字段列表
}
```

**交易模拟失败 (400)**
```json
{
  "error": "simulation_failed",
  "message": "交易不合法，模拟失败",
  "details": "string" // 具体错误信息
}
```

---

### 2. 查询交易状态

**GET** `/api/payments/{transaction_hash}`

根据交易哈希查询交易状态和详情。

#### 路径参数

- `transaction_hash`: 交易哈希值

#### 响应参数

**交易处理中 (200)**
```json
{
  "status": "pending",
  "transaction_hash": "string",
  "message": "交易正在处理中"
}
```

**交易确认成功 (200)**
```json
{
  "status": "confirmed",
  "transaction_hash": "string",
  "success": true,
  "received_amount": "string", // 实际收到的金额
  "currency": "string",        // 货币种类
  "message": "交易已经被区块链确认"
}
```

**交易失败 (200)**
```json
{
  "status": "failed",
  "transaction_hash": "string",
  "success": false,
  "error": "string",
  "message": "交易失败"
}
```

**交易不存在 (404)**
```json
{
  "error": "not_found",
  "message": "交易不存在"
}
```

---

### 3. 查询用户限制

**GET** `/api/users/{user_address}/limits`

根据用户地址查询用户的支付限制信息，包括支付限额、尾部更新次数等。

#### 路径参数

- `user_address`: 用户地址 hex格式

#### 响应参数

**查询成功 (200)**
```json
{
  "code": 1000,
  "data": {
    "user_limits": {
      "payment_limit": number,      // 支付限额
      "tail_update_count": number,  // 尾部更新次数
      "max_tail_updates": number    // 最大尾部更新次数
    }
  }
}
```

**用户地址无效 (400)**
```json
{
  "code": 4000,
  "data": null,
  "message": "Invalid user address format"
}
```

**查询失败 (500)**
```json
{
  "code": 5000,
  "data": null,
  "message": "Failed to get user limits: [具体错误信息]"
}
```

## 使用流程

### 支付流程

1. 前端调用 `POST /api/payments` 创建支付交易
2. 服务器检查字段完整性
3. 服务器进行交易模拟验证
4. 模拟成功则返回交易哈希并提交到区块链，模拟失败则返回错误信息
5. 前端收到交易哈希后，使用 `GET /api/payments/{transaction_hash}` 轮询查询状态
6. 等待状态变为 `confirmed` 或 `failed`

### 用户限制查询流程

1. 前端调用 `GET /api/users/{user_address}/limits` 查询用户限制信息
2. 服务器验证用户地址格式
3. 调用区块链合约的view函数获取用户限制数据
4. 返回包含支付限额、尾部更新次数等信息的响应