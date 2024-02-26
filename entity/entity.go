package entity

import (
	"midtrans/dto"
	"midtrans/model"
)

type MidtransRepository interface {
	Create(data model.Transaction) error
}

type MidtransService interface {
	Create(data model.Transaction) (dto.MidtransResponse, error)
}
