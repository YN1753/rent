package repository

import (
	"context"
	"errors"
	"mime/multipart"
	"rent/internal/model"
	"rent/internal/oss"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HouseRepository struct {
	Mongo *mongo.Database
}

func NewHouseRepository(mongo *mongo.Database) *HouseRepository {
	return &HouseRepository{
		Mongo: mongo,
	}
}

func (h *HouseRepository) Create(house model.House, ctx context.Context) error {
	collection := h.Mongo.Collection("house")
	_, err := collection.InsertOne(ctx, house)
	if err != nil {
		return errors.New("插入失败")
	}
	return nil
}
func (h *HouseRepository) Update(house model.House, ctx context.Context) error {
	collection := h.Mongo.Collection("house")
	_, err := collection.UpdateByID(ctx, house.Id, house)
	if err != nil {
		return errors.New("更新失败")
	}
	return nil
}

func (h *HouseRepository) UploadFile(objectKey string, ctx context.Context, file *multipart.FileHeader, id primitive.ObjectID) error {
	collection := h.Mongo.Collection("house")
	err := oss.UploadFile(objectKey, file)
	if err != nil {
		return err
	}
	count := 0
	for index, s := range objectKey {
		if s == '/' {
			count++
		}
		if count == 2 {
			count = index
			break
		}
	}
	objectKey = objectKey[:count]
	update := bson.M{"$set": bson.M{"file": objectKey}}
	_, err = collection.UpdateOne(ctx, id, update)
	if err != nil {
		return errors.New("更新失败")
	}
	return nil
}
