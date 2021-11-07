package models

import "gorm.io/gorm"

type Wallet struct {
	Uid   int64 `json:"uid"`
	Value int32 `json:"value"`
}

func (Wallet) TableName() string {
	return "wallet"
}


func InsertWallet(DB *gorm.DB, newWallet *Envelope) error {
	return DB.Create(newWallet).Error
}

func UpdateWalletByUid(DB *gorm.DB, uid int64, data *map[string]interface{}) error {
	return DB.Model(&Wallet{}).Where("uid=?", uid).Updates(data).Error
}
