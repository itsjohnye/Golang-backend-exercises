package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) //设置10秒超时时间
	//connect
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx) //disconnect
	//ping
	err = client.Ping(ctx, readpref.Primary()) //ReadPreference读偏好，Primary主节点
	if err != nil {
		log.Fatal(err)
	}
	//ListDatabaseNames
	databases, err := client.ListDatabaseNames(ctx, bson.M{}) //return a []string with each of the database names, bson.M{}是filter argument。
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
}
