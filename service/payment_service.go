package service

import "crowdfunding/model/domain"

type PaymentService interface {
	GetPaymentUrl(transaction domain.Payment, user domain.User) (string, error)
}
