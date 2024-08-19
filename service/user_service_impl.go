package service

import (
	"crowdfunding/helper"
	"crowdfunding/model/domain"
	"crowdfunding/model/web"
	"crowdfunding/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserServiceImplementation struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImplementation{
		UserRepository: userRepository,
	}
}

func (service UserServiceImplementation) Register(request web.UserRequestRegister) (domain.User, error) {
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

	save, err := service.UserRepository.Save(user)
	helper.PanicIfError(err)
	return save, nil
}

func (service UserServiceImplementation) Login(request web.UserRequestLogin) (domain.User, error) {
	//TODO implement me
	email := request.Email
	password := request.Password

	user, err := service.UserRepository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("the email or password you entered is incorrect")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, errors.New("the email or password you entered is incorrect")
	}

	return user, nil
}

func (service UserServiceImplementation) IsEmailAvailable(request web.UserRequestEmailCheck) (bool, error) {
	email := request.Email
	user, err := service.UserRepository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if user.ID == 0 {
		return true, nil
	}
	return false, nil
}
