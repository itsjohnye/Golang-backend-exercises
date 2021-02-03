package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

	database := client.Database("quickstart")
	podcastsCollection := database.Collection("podcasts")
	episodesCollection := database.Collection("episodes")

	//DeleteOne
	result, err := podcastsCollection.DeleteOne(ctx, bson.M{"title": "The Polyglot Developer Podcast"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)

	//DeleteMany
	result, err = episodesCollection.DeleteMany(ctx, bson.M{"duration": 25})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DeleteMany removed %v document(s)\n", result.DeletedCount)

	//Dropping entire collections
	if err = podcastsCollection.Drop(ctx); err != nil {
		log.Fatal(err)
	}
	if err = episodesCollection.Drop(ctx); err != nil {
		log.Fatal(err)
	}
}
