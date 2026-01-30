package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location struct {
	Province     string `json:"province" bson:"province"`
	City         string `json:"city" bson:"city"`
	CityCode     string `json:"cityCode" bson:"cityCode"`
	District     string `json:"district" bson:"district"`
	AdCode       string `json:"adCode" bson:"adCode"`
	ProvinceCode string `json:"provinceCode" bson:"provinceCode"`
}
type House struct {
	Location
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Owner       int                `bson:"owner" json:"owner"`
	Buyer       int                `bson:"buyer" json:"buyer"`
	File        string             `bson:"file" json:"file"`
	Type        string             `bson:"type" json:"type"`
	Position    string             `bson:"position" json:"position"`
	Description string             `bson:"description" json:"description"`
	SellStatus  string             `bson:"sellStatus" json:"sellStatus"`
	Tag         []string           `bson:"tag" json:"tag"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}
type BackHouseInfo struct {
	House
	BackFile []string `json:"backFile"`
}
