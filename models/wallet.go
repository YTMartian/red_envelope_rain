package models

type Wallet struct {
	Uid   int64 `json:"uid"`
	Value int32 `json:"value"`
}

func (Wallet) TableName() string {
	return "wallet"
}
