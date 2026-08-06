package models

type Currency struct {
	ID       int64  `json:"id"`
	Code     string `json:"code"`
	FullName string `json:"fullname"`
	Sign     string `json:"sign"`
}
