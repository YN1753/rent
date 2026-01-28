package model

import "time"

type House struct {
	Owner      int       `json:"owner"`
	Buyer      int       `json:"buyer"`
	File       string    `json:"file"`
	SellStatus string    `json:"sellstatus"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
