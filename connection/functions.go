package connection

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"main.go/config"
)

func Init(conf config.Conf) {
	Database.URI = conf.MongoURI
	Database.Name = conf.MongoDBName
	Database.Timeout = time.Duration(conf.MongoDBTimeout) * time.Second
}

func ConnectDB() {
	opts := options.Client()
	opts.ApplyURI(Database.URI)
	opts.Timeout = &Database.Timeout
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal("error in connecting to mongo uri :", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), Database.Timeout)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error in connecting: ", err)
	}
	log.Println("Connected to MongoDB!")
	MI = MongoInstance{
		Client: client,
		DB:     client.Database(Database.Name),
	}
}
