package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rent/internal/model"
	"rent/internal/service"
	"rent/pkg/common"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{UserService: service.NewUserService()}
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
	var user model.User
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
