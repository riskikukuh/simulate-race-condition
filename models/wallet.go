package models

import "time"

type Wallet struct {
	ID        int64     `gorm:"primary_key;column:id;autoIncrement"`
	UserId    int64     `gorm:"column:user_id"`
	Balance   int64     `gorm:"column:balance"`
	User      *User     `gorm:"foreignKey:user_id;references:id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (u *Wallet) TableName() string {
	return "wallets"
}

type WalletRequest struct {
	UserId  int64
	Balance int64
}
