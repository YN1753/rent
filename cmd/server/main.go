package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rent/internal/config"
	"rent/internal/db"
	"rent/internal/router"
)

func main() {
	config.InitConfig()
	mysqlDB := db.InitMySQL()
	redisDB := db.InitRedis()
	r := gin.Default()
	r.Use(CorsMiddleware())
	router.InitRouter(r, mysqlDB, redisDB)
	r.Run(":8080")
}
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 处理预检请求
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
