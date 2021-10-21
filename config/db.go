package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDBCollection() (*mongo.Collection, error) {

	clientOptions := options.Client().ApplyURI("mongodb+srv://newuser1:Awinner55@cluster0.xpnqk.mongodb.net/goauth?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("goauth").Collection("users")
	return collection, nil
}

func GetContactDBCollection() (*mongo.Collection, error) {

	clientOptions := options.Client().ApplyURI("mongodb+srv://newuser1:Awinner55@cluster0.xpnqk.mongodb.net/goauth?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("goauth").Collection("contact")
	return collection, nil
}
