package wechatapp

type PaocaoAPICount struct {
	UseDate  string `json:"use_date" db:"use_date"`
	UseCount int64  `json:"use_count" db:"use_count"`
}
