package repository

import (
	"midtrans/entity"
	"midtrans/model"

	"gorm.io/gorm"
)

type midtransRepository struct {
	db *gorm.DB
}

func NewMidtransRepository(db *gorm.DB) entity.MidtransRepository {
	return &midtransRepository{
		db: db,
	}
}

// Create implements entity.MidtransRepository.
func (mr *midtransRepository) Create(data model.Transaction) error {
	err := mr.db.Create(&data)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
