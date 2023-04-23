package repository

import (
	"api/src/config"
	"api/src/database"
	"api/src/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type stocks struct {
	client *mongo.Client
}

func AddStock(stock models.Stock) (string, error) {
	client, err := database.ConnectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	collection := client.Database(config.Database).Collection(config.StockCollection)

	result, err := collection.InsertOne(context.Background(), stock)
	if err != nil {
		return "", err
	}

	idStr, _ := result.InsertedID.(primitive.ObjectID)

	return idStr.Hex(), nil
}

func GetStock(name string) (string, error) {
	client, err := database.ConnectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	collection := client.Database(config.Database).Collection(config.StockCollection)

	filter := bson.M{"stock": name}

	var stock models.Stock

	err = collection.FindOne(context.Background(), filter).Decode(&stock)
	if err != nil {
		log.Fatal(err)
	}

	return stock.Stock, nil
}
