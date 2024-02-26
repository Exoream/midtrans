package service

import (
	"errors"
	"midtrans/dto"
	"midtrans/entity"
	"midtrans/model"
	"midtrans/utils"
	"os"

	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type midtransService struct {
	midtransRepository entity.MidtransRepository
}

func NewMidtransService(midtrans entity.MidtransRepository) entity.MidtransService {
	return &midtransService{
		midtransRepository: midtrans,
	}
}

// Create implements entity.MidtransService.
func (ms *midtransService) Create(data model.Transaction) (dto.MidtransResponse, error) {

	godotenv.Load(".env")
	snapClient := &snap.Client{}
	snapClient.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	req := utils.CreateSnapRequest(data)

	res, err := snapClient.CreateTransaction(req)
	if err != nil {
		return dto.MidtransResponse{}, errors.New("error creating midtrans transaction")
	}

	if err := ms.midtransRepository.Create(data); err != nil {
		return dto.MidtransResponse{}, errors.New("error saving transaction to database")
	}

	response := dto.MidtransResponse{
		Token:       res.Token,
		RedirectUrl: res.RedirectURL,
	}

	return response, nil
}
