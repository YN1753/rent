package repository

import (
	"errors"
	"fmt"
	"log"
	"rent/internal/model"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) Register(user model.User) error {
	err := u.DB.Where("username =?", user.Username).First(&model.User{}).Error
	if err == nil {
		return errors.New("账号已存在")
	}
	err = u.DB.Create(&user).Error
	if err != nil {
		log.Println("注册失败", err)
		return errors.New("数据库错误")
	}
	return nil
}

func (u *UserRepository) GetUserInfo(param interface{}) (model.UserInfo, error) {
	var user model.UserInfo
	if email, ok := param.(string); ok {
		err := u.DB.Where("email =?", email).First(&user).Error
		if err != nil {
			return user, err
		}
		return user, err
	} else if id, ok := param.(int); ok {
		err := u.DB.Where("id =?", id).First(&user).Error
		if err != nil {
			return user, err
		}
		return user, err
	}

	return user, errors.New("参数错误")
}

func (u *UserRepository) GetPasswordByUsername(username string) (string, error) {
	var user model.User
	err := u.DB.Where("username =?", username).First(&user).Error
	if err != nil {
		return "", errors.New("账号或密码错误")
	}
	return user.Password, nil
}

func (u *UserRepository) GetPasswordByEmail(email string) (string, error) {
	var user model.User
	err := u.DB.Select("password").Where("email =?", email).First(&user).Error
	if err != nil {
		return "", errors.New("邮箱或密码错误")
	}
	return user.Password, nil
}

func (u *UserRepository) GetPassword(account string) (string, error) {
	err := validator.New().Var(account, "required,email")
	if err != nil {
		return u.GetPasswordByUsername(account)
	} else {
		return u.GetPasswordByEmail(account)
	}
}

func (u *UserRepository) ExistsByUsername(username string) (bool, error) {
	var user model.User
	err := u.DB.Where("username =?", username).First(&user).Error
	if err != nil {
		return false, errors.New("用户名或密码错误")
	}
	return true, nil
}

func (u *UserRepository) ExistsByEmail(email string) (bool, error) {
	var user model.User
	err := u.DB.Where("email =?", email).First(&user).Error
	if err != nil {
		return false, errors.New("邮箱或密码错误")
	}
	return true, nil
}

func (u *UserRepository) Exists(account string) (bool, error) {
	fmt.Println("account:", account)
	err := validator.New().Var(account, "required,email")
	if err != nil {
		return u.ExistsByUsername(account)
	} else {
		return u.ExistsByEmail(account)
	}
}

func (u *UserRepository) GetEmailByUsername(username string) (string, error) {
	var user model.User
	err := u.DB.Where("username =?", username).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.Email, nil
}

func (u *UserRepository) UpdateAvatar(avatar string, id int) error {
	err := u.DB.Table("users").Where("id=?", id).Update("avatar", avatar)
	if err.Error != nil {
		return errors.New("字段更新失败")
	}
	return nil
}
