package model

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string    `json:"username" gorm:"type:varchar(50);uniqueIndex;not null"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Avatar    string    `json:"avatar" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type UserInfo struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
}

func (UserInfo) TableName() string {
	return "users"
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LocationReq struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
