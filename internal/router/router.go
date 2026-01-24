package router

import "github.com/gin-gonic/gin"

func InitRouter(g *gin.Engine) {
	// 初始化路由
	UserRouter(g)
}
