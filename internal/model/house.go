package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type House struct {
	Location
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Owner       int                `bson:"owner" json:"owner"`
	Buyer       int                `bson:"buyer" json:"buyer"`
	File        string             `bson:"file" json:"file"` //房子的图片和视频
	Type        string             `bson:"type" json:"type"` //租赁还是出售
	Area        int                `bson:"area" json:"area"`
	Deposit     int                `bson:"deposit" json:"deposit"`
	HouseType   string             `bson:"houseType" json:"houseType"` //房子类型
	MiniDay     int                `bson:"miniDay" json:"miniDay"`     //最少租赁天数
	Address     string             `bson:"address" json:"address"`     //详细地理位置
	Position    GeoPoint           `bson:"position" json:"position"`   //经纬度
	Description string             `bson:"description" json:"description"`
	SellStatus  string             `bson:"sellStatus" json:"sellStatus"`
	Tag         []string           `bson:"tag" json:"tag"` //房子属性标签
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}
type BackHouseInfo struct {
	House
	BackFile []string `json:"backFile"`
}
