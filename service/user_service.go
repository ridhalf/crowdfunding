package service

import (
	"crowdfunding/model/domain"
	"crowdfunding/model/web"
)

type UserService interface {
	RegisterUser(request web.UserCreateRequest) (domain.User, error)
}
