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

//清楚数据结构，这样能有目标地查询，就不会返回一堆不需要的
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

	// cursor, err := episodesCollection.Find(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var episodes []bson.M
	// if err = cursor.All(ctx, &episodes); err != nil {		//if your expected result set is large, using the *mongo.Cursor.All function might not be the best idea
	// 	log.Fatal(err)
	// }
	// fmt.Println(episodes)

	//用for做一个迭代以便更好地取回数据
	cursor, err := episodesCollection.Find(ctx, bson.M{}) //Find函数有点像ListDatabaseNames函数，bson.M表示无序的一个映射（map of fields）。
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var episode bson.M
		if err = cursor.Decode(&episode); err != nil {
			log.Fatal(err)
		}
		fmt.Println(episode)
	}

	//Reading a Single Document
	var podcast bson.M
	if err = podcastsCollection.FindOne(ctx, bson.M{}).Decode(&podcast); err != nil {
		log.Fatal(err)
	}
	fmt.Println(podcast)

	//Querying Documents from a Collection with a Filter
	filterCursor, err := episodesCollection.Find(ctx, bson.M{"duration": 25})
	if err != nil {
		log.Fatal(err)
	}
	var episodesFiltered []bson.M
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodesFiltered)

	//Sorting Documents in a Query
	opts := options.Find()                       //leverage the FindOptions struct in the MongoDB Go Driver
	opts.SetSort(bson.D{{"duration", -1}}.Map()) //-1表示descending order
	filter := bson.M{"duration": bson.M{"$gt": 24}}
	sortCursor, err := episodesCollection.Find(ctx, filter, opts) //在MongoDB Go Driver v2.0+版本中Find方法要求传入type M map[string]interface{}，直接传bson.D会报错。所以要使用bson.M{}或者bson.D{}.Map()。

	if err != nil {
		log.Fatal(err)
	}
	var episodesSorted []bson.M
	if err = sortCursor.All(ctx, &episodesSorted); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodesSorted)

}
