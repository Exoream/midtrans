package model

import "time"

type Transaction struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	UserId    uint      `json:"user_id"`
	OrderId   string    `json:"order_id"`
	Name      string    `json:"name"`
	Amount    int64     `json:"amount"`
	Qty       int32     `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
