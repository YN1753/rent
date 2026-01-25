package main

import (
	"github.com/gin-gonic/gin"
	"rent/internal/config"
	"rent/internal/db"
	"rent/internal/middleware"
	"rent/internal/router"
)

func main() {
	config.InitConfig()
	mysqlDB := db.InitMySQL()
	redisDB := db.InitRedis()
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	r.Use(gin.Recovery())
	router.InitRouter(r, mysqlDB, redisDB)
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
