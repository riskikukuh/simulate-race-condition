package service

import (
	"simulation-race-condition/models"

	"gorm.io/gorm"
)

type UserServiceImpl struct {
	DB *gorm.DB
}

func NewUserServiceImpl(db *gorm.DB) UserService {
	return &UserServiceImpl{
		DB: db,
	}
}

func (u *UserServiceImpl) Create(data models.UserRequest) (*models.User, error) {
	user := models.User{
		Name:  data.Name,
		Email: data.Email,
	}
	err := u.DB.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserServiceImpl) Update(data *models.User) (*models.User, error) {
	err := u.DB.Transaction(func(tx *gorm.DB) error {
		tx.Updates(map[string]interface{}{
			"name":  data.Name,
			"email": data.Email,
		})
		return nil
	})

	if err != nil {
		return nil, err
	}

	return data, err
}

func (u *UserServiceImpl) FindAll() ([]models.User, error) {
	var result []models.User
	err := u.DB.Find(&result).Error
	if err != nil {
		return make([]models.User, 0), err
	}

	return result, nil
}

func (u *UserServiceImpl) FindById(userId int64) (*models.User, error) {
	var result models.User
	err := u.DB.Take(&result, "id = ?", userId).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (u *UserServiceImpl) DeleteByUserId(userId int64) error {
	// var user models.User
	err := u.DB.Delete(&models.User{}, "id = ?", userId).Error
	return err
}
