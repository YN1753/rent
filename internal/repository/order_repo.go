package repository

import (
	"errors"
	"gorm.io/gorm"
	"rent/internal/model"
)

type OrderRepository struct {
	Mysql *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{Mysql: db}
}

func (o *OrderRepository) Creat(order model.RentOrder) error {
	res := o.Mysql.Create(order)
	if res.Error != nil {
		return errors.New("创建订单失败" + res.Error.Error())
	}
	return nil
}
func (o *OrderRepository) TakeByTenantId(id string) ([]model.RentOrder, error) {
	var orders []model.RentOrder
	if err := o.Mysql.Find(&orders).Where("tenant_id = ?", id).Error; err != nil {
		return nil, errors.New("查询失败" + err.Error())
	}
	return orders, nil
}
