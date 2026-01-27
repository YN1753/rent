package main

import (
	"rent/internal/config"
	"rent/internal/db"
	"rent/internal/middleware"
	"rent/internal/oss"
	"rent/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	mysqlDB := db.InitMySQL()
	redisDB := db.InitRedis()
	oss.InitOSS()
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	r.Use(gin.Recovery())
	router.InitRouter(r, mysqlDB, redisDB)
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
