package db

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"pregen/backend/races"
)

// Replace the placeholder with your Atlas connection string
var (
	cred options.Credential
)

const (
	address = "172.20.0.2:27017"
	dbname  = "data"
)

func connectToDB() *mongo.Client {
	cred.AuthSource = "data"
	cred.Username = "admin"
	cred.Password = "f2h2f342g"
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

func ReadFromDB(dbname, collectionName string) {
	client := connectToDB()

	// begin findOne
	coll := client.Database("data").Collection("races")

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var data bson.D
		if err = cursor.Decode(&data); err != nil {
			log.Fatal(err)
		}
		fmt.Println(data)
	}
}
