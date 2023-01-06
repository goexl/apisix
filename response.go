package apisix

// Response 响应
type Response struct {
	// 键
	Key string `json:"key"`
	// 值
	Value *Value `json:"value"`
}
