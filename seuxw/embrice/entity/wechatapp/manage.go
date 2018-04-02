package wechatapp

// 微信用户信息
type UserInfo struct {
	UserName string `json:"user_name" db:"user_name"` // 用户名称
	UserID   int64  `json:"user_id" db:"user_id"`     //用户ID
}
