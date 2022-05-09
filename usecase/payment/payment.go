package payment

import (
	_entities "capstone/entities"
	"fmt"
	"os"
	"strconv"

	_midtrans "github.com/veritrans/go-midtrans"
)

type service struct {
}

type Service interface {
	GetPaymentURL(transaction _entities.Payment, user _entities.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentURL(transaction _entities.Payment, user _entities.User) (string, error) {
	midclient := _midtrans.NewClient()
	midclient.ServerKey = os.Getenv("SERVER_KEY")
	midclient.ClientKey = os.Getenv("CLIENT_KEY")
	midclient.APIEnvType = _midtrans.Sandbox

	fmt.Println(transaction.ID)

	snapGateway := _midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &_midtrans.SnapReq{
		CustomerDetail: &_midtrans.CustDetail{
			Email: user.Email,
			FName: user.FullName,
		},

		TransactionDetails: _midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(int(transaction.ID)),
			GrossAmt: int64(transaction.TotalPrice),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
