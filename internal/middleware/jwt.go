package middleware

import (
	"rent/internal/config"
	"rent/pkg/common"
	"rent/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
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
		redisClient := redis.NewClient(&redis.Options{
			Addr:     config.Cfg.Redis.Addr,
			Password: config.Cfg.Redis.Password,
			DB:       config.Cfg.Redis.Db,
		})
		// 检查token是否存在于redis中
		if _, err := redisClient.Get(c, "token:"+token).Result(); err != nil {
			common.Success(c, 401, "未登录", nil)
			c.Abort()
			return
		}
		//验证token
		var claims utils.Claims
		secret := config.Cfg.JWT.Secret
		claims, _ = utils.ParseToken(token, secret)
		c.Set("username", claims.Username)
		c.Set("id", claims.Id)
		c.Set("token", token)
		c.Next()
	}
}
