package db

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	address, DBNAME string
	login, passwd   string
)

func InitMongoENV(path string) {
	// Set the file name of the configurations file
	viper.SetConfigName("creds.env")
	// Set the path to look for the configurations file
	viper.AddConfigPath(path)
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	address = viper.GetString("address")
	DBNAME = viper.Get("dbname").(string)
	login = viper.Get("login").(string)
	passwd = viper.Get("passwd").(string)
}

func ConnectToDB() *mongo.Client {
	cred := options.Credential{
		Username: login,
		Password: passwd,
	}
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://" + address + "/" + DBNAME).
		SetServerAPIOptions(serverAPI).SetAuth(cred).SetMaxPoolSize(100000)
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
	if err := client.Database(DBNAME).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
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

	log.Printf("Db version: %+v\n", commandResult["connections"])
}
