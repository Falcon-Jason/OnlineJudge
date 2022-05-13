package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"write_back/pb/judge_result"
)

type writeBackService struct {
	judge_result.UnimplementedResultWriteBackServer
	db *mongo.Database
}

func NewWriteBackService(db *mongo.Database) *writeBackService {
	return &writeBackService{db: db}
}

func (s *writeBackService) WriteBack(
	ctx context.Context, req *judge_result.ResultWriteBackRequest) (
	*judge_result.ResultWriteBackReply, error) {

	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return &judge_result.ResultWriteBackReply{Ok: false}, nil
	}

	_, err = s.db.Collection("submission").
		UpdateOne(ctx, bson.D{
			{"_id", id},
		}, bson.D{
			{"$set", bson.D{{"status", req.Status.String()}}},
		})

	return &judge_result.ResultWriteBackReply{Ok: err == nil}, nil
}

func ConnectMongo(ctx context.Context, uri, dbName string) *mongo.Database {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil || client.Ping(ctx, nil) != nil {
		log.Fatalf("Launch Error: failed to connect to MongoDB with uri %s", uri)
	}

	return client.Database(dbName)
}
