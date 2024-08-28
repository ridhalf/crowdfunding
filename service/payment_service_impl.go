package service

import (
	"crowdfunding/helper"
	"crowdfunding/model/domain"
	"github.com/veritrans/go-midtrans"
	"os"
	"strconv"
)

type PaymentServiceImpl struct {
}

func NewPaymentService() PaymentService {
	return &PaymentServiceImpl{}
}

func (service PaymentServiceImpl) GetPaymentUrl(payment domain.Payment, user domain.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = os.Getenv("SERVER_KEY")
	midclient.ClientKey = os.Getenv("CLIENT_KEY")
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}
	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  helper.ORDER_FORMAT + strconv.Itoa(payment.ID),
			GrossAmt: int64(payment.Amount),
		},
	}
	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}
	return snapTokenResp.RedirectURL, nil
}
