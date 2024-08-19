package service

import (
	"crowdfunding/model/domain"
	"crowdfunding/model/web"
)

type UserService interface {
	RegisterUser(request web.UserRequestRegister) (domain.User, error)
}
