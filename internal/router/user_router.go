package router

import (
	"github.com/gin-gonic/gin"
	"rent/internal/api"
	"rent/internal/middleware"
)

func UserRouter(c *gin.Engine) {
	userHandler := api.NewUserHandler()
	userRouter := c.Group("/user")
	{
		userRouter.POST("/register", userHandler.Register)
		userRouter.POST("/login", userHandler.Login)
	}
	userRouterAuth := userRouter
	{
		userRouterAuth.Use(middleware.AuthMiddleware())
		userRouterAuth.GET("/info", userHandler.GetUserInfo)
	}
}
