package dao

import (
	"flag"
	"log"
)

var (
	mongoUri  = flag.String("mongo_uri", "mongodb://localhost:27017", "The MongoDB server to be connected")
	mongoName = flag.String("mongo_name", "online_judge", "The name of database to be used")
	redisUri  = flag.String("redis_uri", "redis://localhost:6379/", "The Redis server to be connected")
)

func Init() {
	err := connectMongo(*mongoUri, *mongoName)
	if err != nil {
		log.Fatalf("failed to init DAO: %v", err)
	}

	err = connectRedis(*redisUri)
	if err != nil {
		log.Fatalf("failed to init DAO: %v", err)
	}
}

func Close() {
	closeRedis()
	closeMongo()
}
