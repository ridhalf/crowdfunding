package service

import (
	"crowdfunding/helper"
	"crowdfunding/model/domain"
	"crowdfunding/model/web"
	"crowdfunding/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserServiceImpl struct {
	repository repository.UserRepository
}

func NewUserServiceImpl(repository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repository: repository}
}

func (service UserServiceImpl) RegisterUser(request web.UserRequestRegister) (domain.User, error) {
	//TODO implement me
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	helper.PanicIfError(err)

	location, err := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(location).Format("2006-01-02 15:04:05")
	helper.PanicIfError(err)

	user := domain.User{
		Name:         request.Name,
		Email:        request.Email,
		Occupation:   request.Occupation,
		Role:         "user",
		PasswordHash: string(password),
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	save, err := service.repository.Save(user)
	helper.PanicIfError(err)
	return save, nil
}
