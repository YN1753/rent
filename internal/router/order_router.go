package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"rent/internal/api"
)

func OrderRouter(c *gin.Engine, mysql *gorm.DB) {
	OrderHandler := api.NewOrderHandler(mysql)
	OrderRouter := c.Group("/order")
	{
		OrderRouter.POST("/create", OrderHandler.CreateOrder)
	}
}
