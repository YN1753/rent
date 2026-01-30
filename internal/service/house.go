package service

import (
	"context"
	"rent/internal/model"
	"rent/internal/repository"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type HouseService struct {
	HouseRepo *repository.HouseRepository
}

func NewHouseService(mongo *mongo.Database) *HouseService {
	return &HouseService{
		HouseRepo: repository.NewHouseRepository(mongo),
	}
}

func (h *HouseService) Create(house model.House) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := h.HouseRepo.Create(house, ctx)
	if err != nil {
		return err
	}
	return nil
}
