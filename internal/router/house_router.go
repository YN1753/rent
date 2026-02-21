package router

import (
	"rent/internal/api"
	"rent/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func HouseRouter(g *gin.Engine, mdb *mongo.Database, rdb *redis.Client) {
	HouseHandler := api.NewHouseHandler(mdb)
	Houserouter := g.Group("/house")
	Houserouter.Use(middleware.AuthMiddleware(rdb))
	{
		Houserouter.POST("/create", HouseHandler.Create)
		Houserouter.GET("/location/tips", HouseHandler.GetInputTips)
		Houserouter.GET("/location/poi", HouseHandler.GetLocationByPOI)
	}
}
