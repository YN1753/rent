package repository

import (
	"errors"
	"log"
	"rent/internal/model"
	"rent/internal/oss"

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
		signurl, err := oss.CreateSignUrl(user.Avatar)
		if err != nil {
			return user, err
		}
		user.Avatar = signurl[0]
		return user, err
	} else if id, ok := param.(int); ok {
		err := u.DB.Where("id =?", id).First(&user).Error
		signurl, err := oss.CreateSignUrl(user.Avatar)
		if err != nil {
			return user, err
		}
		user.Avatar = signurl[0]
		return user, err
	}

	return user, errors.New("参数错误")
}

func (u *UserRepository) GetPasswordByUsername(username string) (string, error) {
	var user model.User
	err := u.DB.Where("username =?", username).First(&user).Error
	return user.Password, err
}

func (u *UserRepository) GetPasswordByEmail(email string) (string, error) {
	var user model.User
	err := u.DB.Select("password").Where("email =?", email).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.Password, nil
}

func (u *UserRepository) ExistsByUsername(username string) (bool, error) {
	var user model.User
	err := u.DB.Where("username =?", username).First(&user).Error
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (u *UserRepository) ExistsByEmail(email string) (bool, error) {
	var user model.User
	err := u.DB.Where("email =?", email).First(&user).Error
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (u *UserRepository) GetEmailByUsername(username string) (string, error) {
	var user model.User
	err := u.DB.Where("username =?", username).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.Email, nil
}
