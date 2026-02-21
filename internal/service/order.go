package service

import (
	"gorm.io/gorm"
	"rent/internal/model"
	"rent/internal/repository"
)

type OrderService struct {
	OrderRepo *repository.OrderRepository
}

func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{
		OrderRepo: repository.NewOrderRepository(db),
	}
}

func (o *OrderService) Creat(order model.RentOrder) error {
	if err := o.OrderRepo.Creat(order); err != nil {
		return err
	}
	return nil
}
