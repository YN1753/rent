package service

import (
	"errors"
	"fmt"
	"log"
	"rent/internal/config"
	"rent/internal/model"
	"rent/internal/repository"
	"rent/pkg/utils"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		UserRepo: repository.NewUserRepository(),
	}
}

func (u *UserService) Register(user model.User) (error, interface{}) {
	if user.Username == "" || user.Password == "" || user.Phone == "" {
		fmt.Println("用户名、密码、手机号不能为空")
		return errors.New("用户名、密码、手机号不能为空"), nil
	}
	// 检查用户名是否已存在
	existingUser, err := u.UserRepo.ExistsByUsername(user.Username)
	if existingUser {
		fmt.Println("用户名已存在", err)
		return errors.New("用户名已存在"), nil
	}
	// 对密码进行加密
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		fmt.Println("密码加密失败", err)
		return errors.New("密码加密失败"), nil
	}
	user.Password = hashedPassword
	err = u.UserRepo.Register(user)
	if err != nil {
		fmt.Println("数据库插入失败", err)
		return err, nil
	}
	return nil, nil
}
func (u *UserService) Login(user model.User) (error, string) {
	// 检查用户名是否存在
	var token string
	existingUser, _ := u.UserRepo.ExistsByPhone(user.Phone)
	if !existingUser {
		return errors.New("用户名或密码错误"), token
	}
	// 获取用户密码
	hashedPassword, err := u.UserRepo.GetPasswordByPhone(user.Phone)
	fmt.Println(hashedPassword)
	if err != nil {
		return errors.New("获取密码失败"), token
	}
	// 验证密码
	if utils.CheckPassword(user.Password, hashedPassword) {
		fmt.Println("密码错误")
		return errors.New("用户名或密码错误"), token
	}
	// 生成token
	users, err := u.UserRepo.GetUserInfo(user.Phone)
	token, err = utils.GenerateToken(users, config.Cfg.JWT.Secret)
	if err != nil {
		fmt.Println("生成token失败", err)
		return errors.New("生成token失败"), token
	}
	return nil, token
}

func (u *UserService) GetUserInfo(param interface{}) (error, model.UserInfo) {
	user, err := u.UserRepo.GetUserInfo(param)
	log.Println(user)
	if err != nil {
		return err, user
	}
	return nil, user
}
