package wechatapp

type CardBindCount struct {
	CountBind int64   `json:"count_bind" db:"count_bind"`
	CountRate float64 `json:"count_rate" db:"count_rate"`
	CountAll  int64   `json:"count_all"`
}

type CardBindCancel struct {
	Name   string `json:"name" db:"name"`
	CardNO string `json:"card_no" db:"card_no"`
}
