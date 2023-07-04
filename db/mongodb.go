package db

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pregen/backend/backgr"
	"pregen/backend/races"
	"time"
)

var (
	cred options.Credential
)

const (
	address = "172.20.0.2:27017"
	dbname  = "data"
)

func connectToDB() *mongo.Client {

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://" + address + "/" + dbname).SetServerAPIOptions(serverAPI).SetAuth(cred)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	return client
}

func PingMongoDB() {
	var client = connectToDB()
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database(dbname).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	client.Disconnect(context.TODO())

	fmt.Println("[MongoDB] [INFO] Connected: OK")
}

type Book struct {
	Title  string
	Author string
}

func TestInsert() {
	client := connectToDB()
	coll := client.Database(dbname).Collection("races")

	var race races.RaceStruct
	json.Unmarshal(races.RaceHumanJson, &race)

	docs := []interface{}{}

	for _, char := range race {
		fmt.Println(char)
		docs = append(docs, char)
	}

	result, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Documents inserted: %v\n", len(result.InsertedIDs))
	for _, id := range result.InsertedIDs {
		fmt.Printf("Inserted document with _id: %v\n", id)
	}
}

func ReadFromDB(collectionName string) {
	client := connectToDB()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	// begin findOne
	coll := client.Database(dbname).Collection(collectionName)

	var likes []bson.D

	defer client.Disconnect(ctx)

	cursor, err := coll.Find(context.Background(), bson.D{{}})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &likes); err != nil {
		panic(err)
	}
	//fmt.Println(likes)
	var doc []byte
	for _, like := range likes {
		doc, _ = bson.Marshal(like)

	}
	var test backgr.BackBSON
	err = bson.Unmarshal(doc, &test)

	fmt.Println(test)
}
