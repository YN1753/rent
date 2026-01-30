package db

import (
	"context"
	"fmt"
	"log"
	"rent/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Cfg.MongoDB.Url+":"+config.Cfg.MongoDB.Port))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	db := client.Database("users")
	coll := db.Collection("users")
	insertDc := bson.D{
		{Key: "浪凡", Value: "开发者"},
	}
	coll.InsertOne(context.Background(), insertDc)
	var results []bson.M
	cursour, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		log.Println(err)
		return nil
	}
	for cursour.Next(ctx) {
		var result bson.M
		if err := cursour.Decode(&result); err != nil {
			log.Println(err)
		}
		results = append(results, result)
	}
	fmt.Println(results)
	return client
}
