package repository

import (
	"errors"
	"log"
	"rent/internal/db"
	"rent/internal/model"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) Register(user model.User) error {
	err := db.DB.Where("username =?", user.Username).First(&model.User{}).Error
	if err == nil {
		return errors.New("账号已存在")
	}
	err = db.DB.Create(&user).Error
	if err != nil {
		log.Println("注册失败", err)
		return errors.New("数据库错误")
	}
	return nil
}

func (u *UserRepository) GetUserInfo(param interface{}) (model.UserInfo, error) {
	var user model.UserInfo
	if phone, ok := param.(string); ok {
		err := db.DB.Where("phone =?", phone).First(&user).Error
		return user, err
	} else if id, ok := param.(int); ok {
		err := db.DB.Where("id =?", id).First(&user).Error
		return user, err
	}
	return user, errors.New("参数错误")
}

func (u *UserRepository) GetPasswordByUsername(username string) (string, error) {
	var user model.User
	err := db.DB.Where("username =?", username).First(&user).Error
	return user.Password, err
}

func (u *UserRepository) GetPasswordByPhone(phone string) (string, error) {
	var user model.User
	err := db.DB.Select("password").Where("phone =?", phone).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.Password, nil
}

func (u *UserRepository) ExistsByUsername(username string) (bool, error) {
	var user model.User
	err := db.DB.Where("username =?", username).First(&user).Error
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (u *UserRepository) ExistsByPhone(phone string) (bool, error) {
	var user model.User
	err := db.DB.Where("phone =?", phone).First(&user).Error
	if err != nil {
		return false, nil
	}
	return true, nil
}
