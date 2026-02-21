package router

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func InitRouter(g *gin.Engine, db *gorm.DB, rdb *redis.Client, mdb *mongo.Database) {
	// 初始化路由
	UserRouter(g, db, rdb)
	HouseRouter(g, mdb, rdb)
}
