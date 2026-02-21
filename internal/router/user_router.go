package router

import (
	"rent/internal/api"
	"rent/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func UserRouter(c *gin.Engine, db *gorm.DB, rdb *redis.Client) {
	userHandler := api.NewUserHandler(db, rdb)
	userRouter := c.Group("/user")
	{
		userRouter.POST("/register", userHandler.Register)
		userRouter.POST("/login", userHandler.Login)
		userRouter.POST("/gencode", userHandler.GenAuthCode)
		userRouter.POST("/authcode", userHandler.AuthCode)
		userRouter.GET("/location/regeo", userHandler.GetUserLocationByRegeo)
		userRouter.GET("/location/geo", userHandler.GetUserLocationBygeo)
	}
	userRouterAuth := userRouter
	{
		userRouterAuth.Use(middleware.AuthMiddleware(rdb))
		userRouterAuth.GET("/info", userHandler.GetUserInfo)
		userRouterAuth.POST("/uploadavatar", userHandler.UploadAvatar)
		userRouterAuth.POST("/logout", userHandler.Logout)
	}
}
