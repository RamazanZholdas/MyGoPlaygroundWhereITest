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

var (
	mongoURI = "mongodb://localhost:27017"
)

func main() {
	//establish connection
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	if err = client.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	//creating a database in mongo called demo
	demoDB := client.Database("demo")
	if err = demoDB.CreateCollection(ctx, "songs"); err != nil {
		log.Fatal(err)
	}

	//creating a collection in mongo called songs
	songsCollection := demoDB.Collection("songs")
	defer songsCollection.Drop(ctx)

	//inserting one data into collection
	result, err := songsCollection.InsertOne(ctx, bson.D{
		{Key: "name", Value: "Bohemian Rhapsody"},
		{Key: "Duration", Value: time.Now()},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result", result)
	fmt.Println()

	//inserting many data into a collection
	manyResults, err := songsCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{Key: "name", Value: "Crazy Train"},
			{Key: "Duration", Value: time.Now()},
		},
		bson.D{
			{Key: "name", Value: "Give it away"},
			{Key: "Duration", Value: time.Now()},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("manyResult", manyResults)
	fmt.Println()

	//getting all data from collection but its a shit way to do it
	cursor, err := songsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var allSongs []bson.M
	if err = cursor.All(ctx, &allSongs); err != nil {
		log.Fatal(err)
	}

	fmt.Println("All songs:", allSongs)
	fmt.Println()

	//getting all data from collection but it's a good way
	cursor2, err := songsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor2.Close(ctx)

	var singleSong bson.M
	for cursor2.Next(ctx) {
		if err = cursor2.Decode(&singleSong); err != nil {
			log.Fatal(err)
		}
		fmt.Println("song:", singleSong)
	}
	fmt.Println()

	//getting filtered data from mongo
	var songy bson.M
	if err = songsCollection.FindOne(ctx, bson.M{}).Decode(&songy); err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"name": "Give it away"}
	fCursor, err := songsCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	var fSongs []bson.M
	if err = fCursor.All(ctx, &fSongs); err != nil {
		log.Fatal(err)
	}

	fmt.Println("fSongs called give it away", fSongs)
}
