package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	//establish the collection handles
	quickstartDatabase := client.Database("quickstart")
	podcastsCollection := quickstartDatabase.Collection("podcasts")
	episodesCollection := quickstartDatabase.Collection("episodes")

	//InsertOne函数返回两样，一是InsertOneResult和一个error。如果没有错误，InsertOneResult（在例子中是podcastResult变量），有一个InsertedID field（即key：value）.
	podcastResult, err := podcastsCollection.InsertOne(ctx, bson.D{ //bson.D is a Document
		{"title", "The Polyglot Developer Podcast"}, //实则是{Key: "title", Value: "The Polyglot Developer Podcast"},这里可以省去Key-Value notation。
		{"author", "Nic Raboy"},
		{"tags", bson.A{"development", "programming", "coding"}}, //bson.A is an Array
	})

	//InsertMany
	episodeResult, err := episodesCollection.InsertMany(ctx, []interface{}{ //using a slice of interface{}
		bson.D{
			{Key: "podcast", Value: podcastResult.InsertedID}, //前面的InsertedID这里做reference用。可以省略"key:_,value:_"。
			{"title", "GraphQL for API Development"},
			{"description", "Learn about GraphQL from the co-creator of GraphQL, Lee Byron."},
			{"duration", 25},
		},
		bson.D{
			{"podcast", podcastResult.InsertedID},
			{"title", "Progressive Web Application Development"},
			{"description", "Learn about PWA development with Tara Manicsic."},
			{"duration", 32},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted %v documents into episode collection!\n %v\n", len(episodeResult.InsertedIDs), episodeResult.InsertedIDs)

}
