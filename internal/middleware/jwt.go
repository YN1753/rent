package middleware

import (
	"github.com/gin-gonic/gin"
	"rent/internal/config"
	"rent/pkg/common"
	"rent/pkg/utils"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取token
		token := c.GetHeader("Authorization")
		if token == "" {
			common.Success(c, 401, "未登录", nil)
			c.Abort()
			return
		}
		if strings.HasPrefix(token, "Bearer ") {
			token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer "))
		}
		//验证token
		var claims utils.Claims
		secret := config.Cfg.JWT.Secret
		claims, _ = utils.ParseToken(token, secret)
		c.Set("username", claims.Username)
		c.Set("id", claims.Id)
		c.Next()
	}
}
