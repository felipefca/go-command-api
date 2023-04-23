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

type fiat struct {
	client *mongo.Client
}

func AddFiat(stock models.Fiat) (string, error) {
	client, err := database.ConnectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	collection := client.Database(config.Database).Collection(config.FiatCollection)

	result, err := collection.InsertOne(context.Background(), stock)
	if err != nil {
		return "", err
	}

	idStr, _ := result.InsertedID.(primitive.ObjectID)

	return idStr.Hex(), nil
}

func GetFiat(code string) (models.Fiat, error) {
	client, err := database.ConnectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	collection := client.Database(config.Database).Collection(config.FiatCollection)

	filter := bson.D{{Key: "code", Value: code}}

	var fiat models.Fiat
	err = collection.FindOne(context.Background(), filter).Decode(&fiat)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Fiat{}, nil
		}
		panic(err)
	}

	return fiat, nil
}
