package router

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func InitRouter(g *gin.Engine, db *gorm.DB, rdb *redis.Client) {
	// 初始化路由
	UserRouter(g, db, rdb)
}
