package dao

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Mongo *mongo.Database = nil

func connectMongo(uri, dbName string) error {
	if Mongo != nil {
		return errors.New("mongo has been connected")
	}

	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		return err
	}

	Mongo = mongoClient.Database(dbName)
	return nil
}

func closeMongo() {
	if Mongo == nil {
		return
	}

	err := Mongo.Client().Disconnect(context.TODO())
	if err != nil {
		log.Fatalf("failed to close mongo: %v", err)
	}
}

//func InitMongo(ctx context.Context, db *mongo.Database) {
//}
