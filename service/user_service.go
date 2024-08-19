package service

import (
	"crowdfunding/model/domain"
	"crowdfunding/model/web"
)

type UserService interface {
	Register(request web.UserRequestRegister) (domain.User, error)
	Login(request web.UserRequestLogin) (domain.User, error)
}
