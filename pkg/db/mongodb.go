package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	address = "172.20.0.2:27017"
	dbname  = "data"
)

func ConnectToDB() *mongo.Client {

	//переделать чтобы всасывало через вайпер
	cred := options.Credential{
		Username: "admin",
		Password: "f2h2f342g",
	}
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
	var client = ConnectToDB()
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database(dbname).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(client, ctx)

	fmt.Println("[MongoDB] [INFO] Connected: OK")
}

func StatusMongoDB() {
	var client = ConnectToDB()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(client, ctx)
	var commandResult bson.M
	command := bson.D{{"serverStatus", 1}}
	err := client.Database("data").RunCommand(context.TODO(), command).Decode(&commandResult)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Db version: %+v\n", commandResult["connections"])
}

func ReadFromDB(collectionName string) *mongo.Cursor {
	client := ConnectToDB()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	coll := client.Database(dbname).Collection(collectionName)

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(client, ctx)

	return cursor
}
