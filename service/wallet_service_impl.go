package service

import (
	"simulation-race-condition/models"

	"gorm.io/gorm"
)

type WalletServiceImpl struct {
	DB *gorm.DB
}

func NewWalletServiceImpl(db *gorm.DB) WalletService {
	return &WalletServiceImpl{
		DB: db,
	}
}

func (w *WalletServiceImpl) Create(data models.WalletRequest) (*models.Wallet, error) {
	wallet := models.Wallet{
		UserId:  data.UserId,
		Balance: data.Balance,
	}
	err := w.DB.Save(&wallet).Error
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

/*
Without transaction
*/
func (w *WalletServiceImpl) Update(data *models.Wallet) (*models.Wallet, error) {
	// u.DB.Updates(map[string]any{
	// "balance":
	// })

	// w.DB.Save(data)
	err := w.DB.Transaction(func(tx *gorm.DB) error {
		tx.Updates(map[string]interface{}{
			"balance": data.Balance,
			// "email":   data.Email,
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (w *WalletServiceImpl) FindAll() ([]models.Wallet, error) {
	var result []models.Wallet
	err := w.DB.Find(&result).Error
	if err != nil {
		return make([]models.Wallet, 0), err
	}

	return result, nil
}

func (w *WalletServiceImpl) FindById(walletId int64) (*models.Wallet, error) {
	var result models.Wallet
	err := w.DB.Take(&result, "id = ?", walletId).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (w *WalletServiceImpl) DeleteByWalletId(walletId int64) error {
	// var user models.User
	err := w.DB.Delete(&models.Wallet{}, "id = ?", walletId).Error
	return err
}
