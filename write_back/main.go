package main

import (
	"context"
	"flag"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
	"write_back/pb/judge_result"
	"write_back/service"
)

var (
	port      = flag.Int("port", 7102, "The port listened by this service")
	mongoUri  = flag.String("mongo_uri", "mongodb://localhost:27017", "The MongoDB server to be connected")
	mongoName = flag.String("mongo_name", "online_judge", "The name of database to be used")
)

func ConnectMongo(ctx context.Context, uri, dbName string) *mongo.Database {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil || client.Ping(ctx, nil) != nil {
		log.Fatalf("Launch Error: failed to connect to MongoDB with uri %s", uri)
	}

	return client.Database(dbName)
}

func main() {
	flag.Parse()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db := ConnectMongo(ctx, *mongoUri, *mongoName)
	defer db.Client().Disconnect(ctx)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Launch Error: failed to listen to %v", listener.Addr())
	}

	server := grpc.NewServer()
	judge_result.RegisterResultWriteBackServer(server, service.NewWriteBackService(db))
	log.Printf("service listening at %v", listener.Addr())
	log.Printf("with MongoDB connecting to %v", *mongoUri)

	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("Launch Error: failed to start service: %v", err)
	}
}
