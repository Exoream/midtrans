package dto

import "midtrans/model"

type MidtransRequest struct {
	UserId  uint   `json:"user_id"`
	OrderId string `json:"order_id"`
	Name    string `json:"name"`
	Amount  int64  `json:"amount"`
	Qty     int32  `json:"quantity"`
}

func MapMidtransRequestToTransaction(request MidtransRequest) model.Transaction {
	return model.Transaction{
		UserId:  request.UserId,
		OrderId: request.OrderId,
		Name:    request.Name,
		Amount:  request.Amount,
		Qty:     request.Qty,
	}
}
