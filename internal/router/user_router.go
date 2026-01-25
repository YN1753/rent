package router

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"rent/internal/api"
	"rent/internal/middleware"
)

func UserRouter(c *gin.Engine, db *gorm.DB, rdb *redis.Client) {
	userHandler := api.NewUserHandler(db, rdb)
	userRouter := c.Group("/user")
	{
		userRouter.POST("/register", userHandler.Register)
		userRouter.POST("/login", userHandler.Login)
	}
	userRouterAuth := userRouter
	{
		userRouterAuth.Use(middleware.AuthMiddleware())
		userRouterAuth.GET("/info", userHandler.GetUserInfo)
		userRouterAuth.GET("/gencode", userHandler.GenAuthCode)
		userRouterAuth.POST("authcode", userHandler.AuthCode)
	}
}
