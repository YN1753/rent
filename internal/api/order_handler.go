package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"rent/internal/model"
	"rent/internal/service"
	"rent/pkg/common"
)

type OrderHandler struct {
	OrderService *service.OrderService
}

func NewOrderHandler(db *gorm.DB) *OrderHandler {
	return &OrderHandler{
		OrderService: service.NewOrderService(db),
	}
}

func (o *OrderHandler) CreateOrder(c *gin.Context) {
	var order model.RentOrder
	if err := c.ShouldBindJSON(&order); err != nil {
		common.Error(c, 400, "参数错误")
		return
	}
	order.TenantId = c.GetString("id")
	if err := o.OrderService.Creat(order); err != nil {
		common.Error(c, 400, err.Error())
		return
	}
	common.Success(c, 200, "创建成功", nil)
}
