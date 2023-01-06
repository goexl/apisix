package apisix

// Value 值
type Value struct {
	// 内容
	Content string `json:"content"`
	// 修改时间
	UpdateTime int64 `json:"update_time"`
	// 创建时间
	CreateTime int64 `json:"create_time"`
	// 描述
	Description string `json:"desc"`
}
