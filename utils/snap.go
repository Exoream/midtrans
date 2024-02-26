package utils

import (
	"midtrans/model"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func CreateSnapRequest(data model.Transaction) *snap.Request {
	return &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  data.OrderId,
			GrossAmt: data.Amount,
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    data.OrderId,
				Name:  data.Name,
				Price: data.Amount,
				Qty:   1,
			},
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		EnabledPayments: snap.AllSnapPaymentType,
	}
}
