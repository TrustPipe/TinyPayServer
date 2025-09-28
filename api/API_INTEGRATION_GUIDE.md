# TinyPay API Integration Guide

本指南介绍如何使用 OpenAPI 3.0 和 oapi-codegen 为 TinyPay 服务器生成 API 文档和代码。

## 概述

我们使用了 **Design-First** 方法来开发 API：
1. 首先定义 OpenAPI 3.0 规范 (`api/openapi.yaml`)
2. 使用 `oapi-codegen` 生成 Go 代码
3. 实现业务逻辑适配器
4. 提供 Swagger UI 文档界面

## 项目结构

```
api/
├── openapi.yaml        # OpenAPI 3.0 规范定义
├── server.gen.go       # 生成的 Gin 服务器接口
├── types.gen.go        # 生成的数据类型
├── client.gen.go       # 生成的客户端代码
├── spec.gen.go         # 嵌入的规范文件
├── adapter.go          # 业务逻辑适配器
├── docs.go            # 文档服务器
└── oapi-codegen.yaml  # 代码生成配置
```

## 快速开始

### 1. 安装工具

```bash
# 安装 oapi-codegen
make install
# 或者直接安装
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
```

### 2. 生成代码

```bash
# 从 OpenAPI 规范生成所有代码
make generate
```

### 3. 运行服务器

```bash
# 构建并运行
make run

# 或者开发模式运行
make dev
```

### 4. 查看 API 文档

启动服务器后，访问：
- **Swagger UI**: http://localhost:9090/docs
- **OpenAPI 规范**: http://localhost:9090/api/openapi.yaml
- **健康检查**: http://localhost:9090/api/health

## OpenAPI 规范特性

### 1. 完整的数据模型 <mcreference link="https://github.com/oapi-codegen/oapi-codegen" index="1">1</mcreference>

我们在 `components/schemas` 中定义了所有数据结构：

```yaml
components:
  schemas:
    PaymentRequest:
      type: object
      required:
        - payer_addr
        - otp
        - payee_addr
        - amount
      properties:
        payer_addr:
          type: string
          pattern: '^0x[a-fA-F0-9]{1,64}$'
        # ... 其他字段
```

### 2. 详细的响应定义 <mcreference link="https://www.networknt.com/development/best-practices/openapi3/" index="3">3</mcreference>

每个端点都有完整的响应定义和示例：

```yaml
responses:
  '200':
    description: 交易模拟成功，已提交到区块链
    content:
      application/json:
        schema:
          $ref: '#/components/schemas/PaymentSuccessResponse'
        examples:
          success:
            summary: 成功提交交易
            value:
              status: "submitted"
              transaction_hash: "0x1a2b..."
```

### 3. 输入验证 <mcreference link="https://spec.openapis.org/oas/v3.0.3.html" index="2">2</mcreference>

使用 JSON Schema 验证：
- 必需字段检查
- 数据类型验证
- 正则表达式模式匹配
- 枚举值限制

## 代码生成详解

### 1. 服务器接口生成

```bash
oapi-codegen -package=api -generate gin api/openapi.yaml > api/server.gen.go
```

生成的接口：
```go
type ServerInterface interface {
    HealthCheck(c *gin.Context)
    CreatePayment(c *gin.Context)
    GetTransactionStatus(c *gin.Context, transactionHash string)
}
```

### 2. 类型定义生成 <mcreference link="https://ldej.nl/post/generating-go-from-openapi-3/" index="4">4</mcreference>

```bash
oapi-codegen -package=api -generate types api/openapi.yaml > api/types.gen.go
```

生成强类型的 Go 结构体，包括：
- 请求/响应模型
- 枚举常量
- JSON 标签

### 3. 客户端代码生成

```bash
oapi-codegen -package=api -generate client api/openapi.yaml > api/client.gen.go
```

生成的客户端可以直接使用：
```go
client, err := api.NewClient("http://localhost:9090")
resp, err := client.CreatePayment(ctx, api.PaymentRequest{...})
```

## 适配器模式

我们使用适配器模式连接生成的接口和现有业务逻辑：

```go
// api/adapter.go
type APIServer struct {
    aptosClient *client.AptosClient
}

func (s *APIServer) CreatePayment(c *gin.Context) {
    // 1. 解析请求
    var req PaymentRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        // 处理错误
    }

    // 2. 调用业务逻辑
    txHash, err := s.aptosClient.SubmitPayment(...)

    // 3. 返回标准响应
    c.JSON(http.StatusOK, PaymentSuccessResponse{...})
}
```

## 最佳实践 <mcreference link="https://dev.to/nikita_rykhlov/go-tools-code-generation-from-openapi-specs-in-go-with-oapi-codegen-3jc1" index="5">5</mcreference>

### 1. 规范设计

- **使用语义化的 operationId**: 生成的方法名更清晰
- **定义完整的 components**: 避免重复定义
- **添加详细的描述和示例**: 提高文档质量
- **使用适当的 HTTP 状态码**: 符合 RESTful 规范

### 2. 代码组织

- **分离生成代码和业务代码**: 避免覆盖自定义逻辑
- **使用适配器模式**: 连接生成接口和现有代码
- **版本控制生成的代码**: 便于跟踪变更

### 3. 开发流程

1. 修改 `api/openapi.yaml`
2. 运行 `make generate` 重新生成代码
3. 更新适配器实现（如需要）
4. 测试 API 功能
5. 检查文档更新

## 文档功能

### Swagger UI

访问 http://localhost:9090/docs 可以：
- 浏览所有 API 端点
- 查看请求/响应模型
- 在线测试 API
- 下载 OpenAPI 规范

### 自动化文档

文档会自动从 OpenAPI 规范生成，包括：
- API 端点列表
- 参数说明
- 响应格式
- 错误代码
- 使用示例

## 扩展 API

### 添加新端点

1. 在 `api/openapi.yaml` 中添加新路径：

```yaml
paths:
  /api/new-endpoint:
    post:
      summary: 新端点
      operationId: newEndpoint
      # ... 其他定义
```

2. 重新生成代码：

```bash
make generate
```

3. 在适配器中实现新方法：

```go
func (s *APIServer) NewEndpoint(c *gin.Context) {
    // 实现逻辑
}
```

### 修改现有端点

1. 更新 OpenAPI 规范
2. 重新生成代码
3. 更新适配器实现
4. 测试兼容性

## 故障排除

### 常见问题

1. **生成代码编译错误**
   - 检查 OpenAPI 规范语法
   - 确保所有引用的 schema 都已定义

2. **路由冲突**
   - 检查路径参数定义
   - 确保 operationId 唯一

3. **类型不匹配**
   - 检查 JSON 标签
   - 确认数据类型定义正确

### 调试技巧

```bash
# 验证 OpenAPI 规范
make validate

# 查看生成的代码
cat api/server.gen.go

# 检查服务器日志
make dev
```

## 总结

使用 oapi-codegen 的优势：

1. **类型安全**: 编译时检查 API 契约
2. **自动化**: 减少手动编写样板代码
3. **一致性**: 确保实现与规范一致
4. **文档**: 自动生成交互式文档
5. **客户端**: 自动生成类型安全的客户端

这种 Design-First 方法确保了 API 的质量和一致性，同时提高了开发效率。