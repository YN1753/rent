package api

import (
	"fmt"
	"rent/internal/model"
	"rent/internal/service"
	"rent/pkg/common"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(db *gorm.DB, rdb *redis.Client) *UserHandler {
	return &UserHandler{UserService: service.NewUserService(db, rdb)}
}

func (u *UserHandler) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		common.Error(c, 400, "参数错误")
		return
	}
	fmt.Println(user)
	if err, _ := u.UserService.Register(user); err != nil {
		common.Error(c, 500, err.Error())
		return
	} else {
		common.Success(c, 200, "注册成功", nil)
	}
}

func (u *UserHandler) Login(c *gin.Context) {
	var user model.LoginRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		common.Error(c, 400, "参数错误")
		return
	}
	var token string
	var err error
	if err, token = u.UserService.Login(user); err != nil {
		common.Error(c, 500, err.Error())
		return
	}
	common.Success(c, 200, "登录成功", gin.H{
		"token": token,
	})
}

func (u *UserHandler) Logout(c *gin.Context) {
	token, _ := c.Get("token")
	if err := u.UserService.Logout(token.(string)); err != nil {
		common.Error(c, 500, err.Error())
		return
	}
	common.Success(c, 200, "退出登录成功", nil)
}

func (u *UserHandler) GetUserInfo(c *gin.Context) {
	id, _ := c.Get("id")
	var user model.UserInfo
	var err error
	if err, user = u.UserService.GetUserInfo(id); err != nil {
		fmt.Println("用户信息：", user)
		common.Error(c, 500, err.Error())
		return
	}
	common.Success(c, 200, "获取成功", user)
}

func (u *UserHandler) GenAuthCode(c *gin.Context) {
	type Email struct {
		Email string `json:"email"`
	}
	var email Email
	if err := c.ShouldBindJSON(&email); err != nil {
		common.Error(c, 400, "参数错误")
		return
	}
	err := u.UserService.GenCode(c, email.Email)
	if err != nil {
		common.Error(c, 400, err.Error())
		return
	}
	common.Success(c, 200, "发送成功", nil)
}

func (u *UserHandler) AuthCode(c *gin.Context) {
	type Code struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}
	var code Code
	if err := c.ShouldBindJSON(&code); err != nil {
		common.Error(c, 400, "参数错误")
		return
	}
	err := u.UserService.AuthCode(code.Code, code.Email)
	if err != nil {
		common.Error(c, 400, "验证码错误")
		return
	}
	common.Success(c, 200, "验证成功", nil)
}

func (u *UserHandler) UploadAvatar(c *gin.Context) {
	err := u.UserService.UploadAvatar(c)
	if err != nil {
		common.Error(c, 500, err.Error())
		return
	}
	common.Success(c, 200, "上传头像成功", nil)
}

func (u *UserHandler) GetLocation(c *gin.Context) {

}
