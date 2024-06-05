package application

import (
	"context"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
)

type PaymentReadAppSvc struct{}

func NewPaymentReadAppSvc() *PaymentReadAppSvc {
	return &PaymentReadAppSvc{}
}

// 根据id来查询读者信息
func (receiver PaymentReadAppSvc) QueryById(ctx context.Context, paymentId int64) (*do.PaymentDO, error) {
	payment, err := repo.NewPaymentRepo().QueryById(ctx, paymentId)
	if err != nil {
		return nil, err
	}
	return payment, nil
}
