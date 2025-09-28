package api

// 业务状态码定义
const (
	// 成功状态码 (1000-1999)
	CodeServerHealthy        = 1000 // 服务器运行正常
	CodeTransactionCreated   = 1001 // 交易创建成功
	CodeTransactionPending   = 1002 // 交易处理中
	CodeTransactionConfirmed = 1003 // 交易确认成功

	// 错误状态码 (2000-2999)
	CodeAmountMustBePositive = 2000 // 金额必须大于0
	CodeAmountExceedsLimit   = 2001 // 金额超出限制
	CodeInsufficientBalance  = 2002 // 余额不足
	CodeInvalidOpt           = 2003 // OPT 不正确
	CodeMissingFields        = 2004 // 缺少必需字段
	CodeTransactionNotFound  = 2005 // 交易不存在
	CodeInvalidCurrency      = 2006 // 无效的货币种类
)

// CreateApiResponse 创建统一的API响应
func CreateApiResponse(code int, data interface{}) ApiResponse {
	var dataMap *map[string]interface{}
	if data != nil {
		if m, ok := data.(map[string]interface{}); ok {
			dataMap = &m
		} else {
			// 如果不是map类型，转换为map
			dataMap = &map[string]interface{}{"result": data}
		}
	}
	return ApiResponse{
		Code: code,
		Data: dataMap,
	}
}

// CreateApiResponseWithMap 创建带有map数据的API响应
func CreateApiResponseWithMap(code int, data map[string]interface{}) ApiResponse {
	return ApiResponse{
		Code: code,
		Data: &data,
	}
}

// CreateApiResponseWithNullData 创建data为null的API响应
func CreateApiResponseWithNullData(code int) ApiResponse {
	return ApiResponse{
		Code: code,
		Data: nil,
	}
}
