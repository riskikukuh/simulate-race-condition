package models

import "time"

type User struct {
	ID        int64     `gorm:"primary_key;column:id;autoIncrement"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email"`
	Wallet    Wallet    `gorm:"foreignKey:user_id;references:id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (u *User) TableName() string {
	return "users"
}

type UserRequest struct {
	Name  string
	Email string
}
