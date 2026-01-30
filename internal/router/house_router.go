package router

import (
	"rent/internal/api"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func HouseRouter(g *gin.Engine, mdb *mongo.Database) {
	HouseHandler := api.NewHouseHandler(mdb)
	Houserouter := g.Group("/house")
	{
		Houserouter.POST("/create", HouseHandler.Create)
	}
}
