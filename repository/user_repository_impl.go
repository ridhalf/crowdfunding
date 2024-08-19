package repository

import (
	"crowdfunding/helper"
	"crowdfunding/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

func (repository UserRepositoryImpl) Save(user domain.User) (domain.User, error) {
	//TODO implement me
	err := repository.db.Create(&user).Error
	helper.PanicIfError(err)
	return user, nil
}
