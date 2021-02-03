package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
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

	podcastsCollection := client.Database("quickstart").Collection("podcasts")
	id, _ := primitive.ObjectIDFromHex("60196282d4c275bbebdded99") // a hexadecimal value that you've received from a RESTful API request

	//UpdateOne
	result, err := podcastsCollection.UpdateOne( //UpdateOne function returns UpdateResult(as 'result' variable here) and and error.
		ctx,               //context
		bson.M{"_id": id}, //filter
		bson.M{"$set": bson.M{"author": "NNNNNic RRRRRaboy"}}, //update
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v Documents, %v Upserted\n", result.ModifiedCount, result.UpsertedCount)

	//UpdateMany
	result, err = podcastsCollection.UpdateMany(
		ctx,
		bson.M{"title": "The Polyglot Developer Podcast"},
		bson.M{"$set": bson.M{"author": "Nicolas Raboy"}},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)

	//ReplaceOne
	result, err = podcastsCollection.ReplaceOne(
		ctx,
		bson.M{"author": "Nicolas Raboy"}, //When working with the ReplaceOne function, update operators such as $set cannot be used since it is a complete replace rather than an update of particular fields.
		bson.M{
			"title":  "The AAABBBCCC Show",
			"author": "Nicolas Raboy",
		},
	)
	fmt.Printf("Replaced %v Documents!\n", result.ModifiedCount)

}
