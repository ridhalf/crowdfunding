package repository

import (
	"crowdfunding/helper"
	"crowdfunding/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (repository UserRepositoryImpl) Save(user domain.User) (domain.User, error) {
	//TODO implement me
	err := repository.db.Create(&user).Error
	helper.PanicIfError(err)
	return user, nil
}

func (repository UserRepositoryImpl) FindByEmail(email string) (domain.User, error) {
	user := domain.User{}
	err := repository.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
