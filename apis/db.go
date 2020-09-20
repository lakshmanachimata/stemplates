package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClinent mongo.Client

func ConnectDB() {
	var configdata ConfigData = readConfig("config.json")

	var dbString string = "mongodb://" + configdata.DBIp + ":" + configdata.DBPort

	dbClinent, err := mongo.NewClient(options.Client().ApplyURI(dbString))
	if err != nil {
		log.Fatal(err)
	}
	var ctx, _ = context.WithTimeout(context.Background(), time.Duration(configdata.dbTimeOut)*time.Second)
	err = dbClinent.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = dbClinent.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB Creatd.")

}

func getDB(dbname string) *mongo.Database {
	return dbClinent.Database(dbname)
}

func getCollection(db mongo.Database, collection string) *mongo.Collection {
	return db.Collection(collection)
}

func DisConnectDB() {
	var err = dbClinent.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
