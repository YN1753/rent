package service

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"rent/internal/config"
	"rent/internal/model"
	"rent/internal/oss"
	"rent/internal/repository"
	"rent/pkg/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserService struct {
	UserRepo *repository.UserRepository
	Redis    *redis.Client
}

func NewUserService(db *gorm.DB, rdb *redis.Client) *UserService {
	return &UserService{
		UserRepo: repository.NewUserRepository(db),
		Redis:    rdb,
	}
}

func (u *UserService) Register(user model.User) (error, interface{}) {
	fmt.Println("注册用户", user.Username)
	if user.Username == "" || user.Password == "" || user.Email == "" {
		log.Println("用户名、密码、邮箱不能为空")
		return errors.New("用户名、密码、邮箱不能为空"), nil
	}
	// 检查邮箱是否已存在
	existingEmail, err := u.UserRepo.ExistsByEmail(user.Email)
	if existingEmail {
		log.Println("邮箱已存在", err)
		return errors.New("邮箱已存在"), nil
	}
	// 对密码进行加密
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println("密码加密失败", err)
		return errors.New("密码加密失败"), nil
	}
	user.Password = hashedPassword
	err = u.UserRepo.Register(user)
	if err != nil {
		log.Println("数据库插入失败", err)
		return err, nil
	}
	return nil, nil
}
func (u *UserService) Login(user model.LoginRequest) (error, string) {
	// 检查用户名是否存在
	var token string
	existingUser, err := u.UserRepo.Exists(user.Account)
	if !existingUser {
		return err, token
	}
	// 获取用户密码
	hashedPassword, err := u.UserRepo.GetPassword(user.Account)
	if err != nil {
		log.Println("获取密码失败", err)
		return err, token
	}
	// 验证密码
	if utils.CheckPassword(user.Password, hashedPassword) {
		return errors.New("账号或密码错误"), token
	}
	// 生成token
	users, err := u.UserRepo.GetUserInfo(user.Account)
	token, err = utils.GenerateToken(users, config.Cfg.JWT.Secret)
	if err != nil {
		log.Println("生成token失败", err)
		return errors.New("生成token失败"), token
	}
	u.Redis.Set(context.Background(), "token:"+token, token, 24*time.Hour)
	return nil, token
}

func (u *UserService) Logout(token string) error {
	err := u.Redis.Del(context.Background(), "token:"+token).Err()
	if err != nil {
		log.Println("删除token失败", err)
		return err
	}
	return nil
}

func (u *UserService) GetUserInfo(account interface{}) (error, model.UserInfo) {
	user, err := u.UserRepo.GetUserInfo(account)
	log.Println(user)
	if err != nil {
		log.Println("获取用户信息失败", err)
		return err, user
	}
	signurl, err := oss.CreateSignUrl(user.Avatar)
	if err != nil {
		return err, user
	}
	user.Avatar = signurl[0]
	return nil, user
}

func (u *UserService) AuthCode(code string, email string) error {
	ctx := context.Background()
	result, err := u.Redis.Get(ctx, email).Result()
	if err != nil {
		fmt.Println(err)
		return err
	}
	if code == result {
		u.Redis.Del(ctx, email)
		return nil
	} else {
		return errors.New("验证码错误")
	}
}

func (u *UserService) GenCode(c *gin.Context, email string) error {
	// username, ok := c.Get("username")
	// if !ok {
	// 	return errors.New("获取用户名失败")
	// }
	// emailExist, err := u.UserRepo.GetEmailByUsername(username.(string))
	// if err != nil {
	// 	return errors.New("获取邮箱失败")
	// }
	// if emailExist != email {
	// 	return errors.New("邮箱错误")
	// }
	charset := "0123456789"
	code := make([]byte, 6)
	n, err := rand.Read(code)
	if err != nil || n != 6 {
		return errors.New("生成验证码失败")
	}
	for i := range code {
		code[i] = charset[code[i]%byte(len(charset))]
	}
	ctx := context.Background()
	err = utils.SendEmail(email, "验证码", "这是你的验证码"+string(code)+",有效期为5分钟")
	if err != nil {
		return errors.New("发送验证码失败")
	}
	u.Redis.Set(ctx, email, string(code), 300*time.Second)
	return nil
}

func (u *UserService) UploadAvatar(c *gin.Context) error {
	id := c.GetInt("id")
	name := c.GetString("username")
	file, err := c.FormFile("img")
	ext := filepath.Ext(file.Filename)
	objectKey := "avatar/" + name + strconv.Itoa(id) + ext
	if err != nil {
		return errors.New("获取图片失败")
	}
	err = oss.UploadFile(objectKey, file)
	if err != nil {
		return err
	}
	err = u.UserRepo.UpdateAvatar(objectKey, id)
	return err
}
