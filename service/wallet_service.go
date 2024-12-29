package service

import (
	"simulation-race-condition/models"
)

type WalletService interface {
	Create(data models.WalletRequest) (*models.Wallet, error)
	Update(data *models.Wallet) (*models.Wallet, error)
	FindAll() ([]models.Wallet, error)
	FindById(walletId int64) (*models.Wallet, error)
	DeleteByWalletId(walletId int64) error
}
