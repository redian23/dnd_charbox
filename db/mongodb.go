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

var JsonTemp = []byte(`[
   {
      "charReq":[
         [
            "Intelligence",
            "BodyDifficulty",
            "Charisma"
         ],
         [
            "Intelligence",
            "Dexterity",
            "Charisma"
         ],
         [
            "Intelligence",
            "Wisdom",
            "Charisma"
         ]
      ],
      "hits":{
         "hit_dice":"1к6",
         "hit_count":6
      },
      "saving_throws":[
         "Intelligence",
         "Wisdom"
      ],
      "skills":{
         "random_count":2,
         "skill_list":[
            "History",
            "Arcana",
            "Medicine",
            "Insight",
            "Investigation",
            "Religion"
         ]
      },
      "className":"Wizard",
      "classNameRU":"Волшебник",
      "background":[
         "Guild artisan",
         "Sage",
         "Soldier",
         "Folk hero"
      ]
   },
   {
      "charReq":[
         [
            "Intelligence",
            "Dexterity",
            "BodyDifficulty"
         ],
         [
            "Intelligence",
            "Dexterity",
            "Wisdom"
         ]
      ],
      "hits":{
         "hit_dice":"1к6",
         "hit_count":6
      },
      "saving_throws":[
         "Dexterity",
         "Intelligence"
      ],
      "skills":{
         "random_count":"3",
         "skill_list":[
            "Perception",
            "Survival",
            "History",
            "Sleight Of Hand",
            "Arcana",
            "Medicine",
            "Nature",
            "Insight"
         ]
      },
      "className":"Alchemist",
      "classNameRU":"Алхимик",
      "background":[
         "Noble",
         "Guild artisan",
         "Soldier"
      ]
   },
   {
      "charReq":[
         [
            "Strength",
            "Dexterity",
            "Intelligence"
         ],
         [
            "Strength",
            "Dexterity",
            "BodyDifficulty"
         ],
         [
            "Strength",
            "Dexterity",
            "Charisma"
         ]
      ],
      "hits":{
         "hit_dice":"1к10",
         "hit_count":"10"
      },
      "saving_throws":[
         "Strength",
         "Body Difficulty"
      ],
      "skills":{
         "random_count":2,
         "skill_list":[
            "Acrobatics",
            "Athletics",
            "Perception",
            "Survival",
            "Intimidation",
            "History",
            "Insight",
            "Animal Handling"
         ]
      },
      "className":"Fighter",
      "classNameRU":"Воин",
      "background":[
         "Urchin",
         "Noble",
         "Sailor",
         "Knight",
         "Folk hero",
         "Soldier",
         "Outlander"
      ]
   },
   {
      "charReq":[
         [
            "Strength",
            "BodyDifficulty",
            "Wisdom"
         ],
         [
            "Strength",
            "BodyDifficulty",
            "Charisma"
         ],
         [
            "Strength",
            "BodyDifficulty",
            "Intelligence"
         ]
      ],
      "hits":{
         "hit_dice":"1к12",
         "hit_count":12
      },
      "saving_throws":[
         "Strength",
         "Body Difficulty"
      ],
      "skills":{
         "random_count":2,
         "skill_list":[
            "Athletics",
            "Perception",
            "Survival",
            "Intimidation",
            "Nature",
            "Animal Handling"
         ]
      },
      "className":"Barbarian",
      "classNameRU":"Варвар",
      "background":[
         "Sailor",
         "Folk hero",
         "Pirate",
         "Soldier",
         "Outlander"
      ]
   },
   {
      "charReq":[
         [
            "Strength",
            "Charisma",
            "BodyDifficulty"
         ],
         [
            "Strength",
            "Charisma",
            "Wisdom"
         ]
      ],
      "hits":{
         "hit_dice":"1к10",
         "hit_count":10
      },
      "saving_throws":[
         "Wisdom",
         "Charisma"
      ],
      "skills":{
         "random_count":2,
         "skill_list":[
            "Athletics",
            "Intimidation",
            "Medicine",
            "Insight",
            "Religion",
            "Persuasion"
         ]
      },
      "className":"Paladin",
      "classNameRU":"Паладин",
      "background":[
         "Noble",
         "Hermit",
         "Soldier"
      ]
   },
   {
      "charReq":[
         [
            "Dexterity",
            "Wisdom",
            "Strength"
         ],
         [
            "Dexterity",
            "BodyDifficulty",
            "Charisma"
         ],
         [
            "Dexterity",
            "BodyDifficulty",
            "Wisdom"
         ]
      ],
      "hits":{
         "hit_dice":"1к8",
         "hit_count":8
      },
      "saving_throws":[
         "Strength",
         "Dexterity"
      ],
      "skills":{
         "random_count":2,
         "skill_list":[
            "Acrobatics",
            "Athletics",
            "History",
            "Insight",
            "Religion",
            "Stealth"
         ]
      },
      "className":"Monk",
      "classNameRU":"Монах",
      "background":[
         "Folk hero",
         "Hermit",
         "Acolyte"
      ]
   },
   {
      "charReq":[
         [
            "Dexterity",
            "Intelligence",
            "Charisma"
         ],
         [
            "Dexterity",
            "Intelligence",
            "Wisdom"
         ]
      ],
      "hits":{
         "hit_dice":"1к8",
         "hit_count":8
      },
      "saving_throws":[
         "Dexterity",
         "Intelligence"
      ],
      "skills":{
         "random_count":4,
         "skill_list":[
            "Acrobatics",
            "Athleticism",
            "Perception",
            "Performance",
            "Intimidation",
            "Sleight Of Hand",
            "Deception",
            "Insight",
            "Investigation",
            "Stealth",
            "Persuasion"
         ]
      },
      "className":"Rogue",
      "classNameRU":"Разбойник",
      "background":[
         "Urchin",
         "Pirate",
         "Criminal",
         "Charlatan",
         "Outlander"
      ]
   },
   {
      "charReq":[
         [
            "Dexterity",
            "Wisdom",
            "Charisma"
         ],
         [
            "Strength",
            "Wisdom",
            "Dexterity"
         ]
      ],
      "hits":{
         "hit_dice":"1к10",
         "hit_count":10
      },
      "saving_throws":[
         "Strength",
         "Dexterity"
      ],
      "skills":{
         "random_count":3,
         "skill_list":[
            "Athleticism",
            "Perception",
            "Survival",
            "Nature",
            "Insight",
            "Investigation",
            "Stealth",
            "Animal Handling"
         ]
      },
      "className":"Ranger",
      "classNameRU":"Лучник",
      "background":[
         "Guild artisan",
         "Folk hero",
         "Criminal",
         "Outlander"
      ]
   },
   {
      "charReq":[
         [
            "Wisdom",
            "BodyDifficulty",
            "Dexterity"
         ],
         [
            "Wisdom",
            "BodyDifficulty",
            "Strength"
         ],
         [
            "Wisdom",
            "BodyDifficulty",
            "Intelligence"
         ]
      ],
      "hits":{
         "hit_dice":"1к8",
         "hit_count":8
      },
      "saving_throws":[
         "Intelligence",
         "Wisdom"
      ],
      "skills":{
         "random_count":2,
         "skill_list":[
            "Perception",
            "Survival",
            "Magic",
            "Medicine",
            "Animal Handling",
            "Nature",
            "Insight",
            "Religion"
         ]
      },
      "className":"Druid",
      "classNameRU":"Друид",
      "background":[
         "Guild artisan",
         "Folk hero",
         "Hermit",
         "Acolyte",
         "Outlander"
      ]
   },
   {
      "charReq":[
         [
            "Wisdom",
            "BodyDifficulty",
            "Charisma"
         ],
         [
            "Wisdom",
            "Strength",
            "Charisma"
         ]
      ],
      "hits":{
         "hit_dice":"1к8",
         "hit_count":8
      },
      "saving_throws":[
         "Wisdom",
         "Charisma"
      ],
      "skills":{
         "random_count":2,
         "skill_list":[
            "History",
            "Medicine",
            "Insight",
            "Religion",
            "Persuasion"
         ]
      },
      "className":"Cleric",
      "classNameRU":"Клирик",
      "background":[
         "Sage",
         "Folk hero",
         "Hermit",
         "Acolyte",
         "Charlatan"
      ]
   },
   {
      "charReq":[
         [
            "Charisma",
            "BodyDifficulty",
            "Wisdom"
         ]
      ],
      "hits":{
         "hit_dice":"1к8",
         "hit_count":8
      },
      "saving_throws":[
         "Wisdom",
         "Charisma"
      ],
      "skills":{
         "random_count":"2",
         "skill_list":[
            "Intimidation",
            "History",
            "Arcana",
            "Deception",
            "Nature",
            "Investigation",
            "Religion"
         ]
      },
      "className":"Warlock",
      "classNameRU":"Чернокнижник",
      "background":[
         "Hermit",
         "Sage",
         "Charlatan"
      ]
   },
   {
      "charReq":[
         [
            "Charisma",
            "BodyDifficulty",
            "Intelligence"
         ]
      ],
      "hits":{
         "hit_dice":"1к6",
         "hit_count":6
      },
      "saving_throws":[
         "Body Difficulty",
         "Charisma"
      ],
      "skills":{
         "random_count":2,
         "skill_list":[
            "Intimidation",
            "Arcana",
            "Deception",
            "Insight",
            "Religion",
            "Persuasion"
         ]
      },
      "className":"Sorcerer",
      "classNameRU":"Заклинатель",
      "background":[
         "Guild artisan",
         "Sage",
         "Folk hero",
         "Hermit"
      ]
   },
   {
      "charReq":[
         [
            "Charisma",
            "Dexterity",
            "Wisdom"
         ],
         [
            "Charisma",
            "Dexterity",
            "BodyDifficulty"
         ],
         [
            "Charisma",
            "Dexterity",
            "Intelligence"
         ]
      ],
      "hits":{
         "hit_dice":"1к8",
         "hit_count":8
      },
      "saving_throws":[
         "Dexterity",
         "Charisma"
      ],
      "skills":{
         "random_count":3,
         "skill_list":[
            "Acrobatics",
            "AnimalHandling",
            "Arcana",
            "Athletics",
            "Deception",
            "History",
            "Insight",
            "Intimidation",
            "Investigation",
            "Medicine",
            "Nature",
            "Perception",
            "Performance",
            "Persuasion",
            "Religion",
            "Sleight Of Hand",
            "Stealth",
            "Survival"
         ]
      },
      "className":"Bard",
      "classNameRU":"Бард",
      "background":[
         "Artist",
         "Noble",
         "Urchin",
         "Charlatan"
      ]
   },
   {
      "charReq":[
         [
            "Intelligence",
            "BodyDifficulty",
            "Wisdom"
         ],
         [
            "Intelligence",
            "BodyDifficulty",
            "Dexterity"
         ]
      ],
      "hits":{
         "hit_dice":"1к8",
         "hit_count":8
      },
      "saving_throws":[
         "Body Difficulty",
         "Intelligence"
      ],
      "skills":{
         "random_count":2,
         "skill_list":[
            "Perception",
            "History",
            "Sleight Of Hand",
            "Arcana",
            "Medicine",
            "Nature",
            "Investigation"
         ]
      },
      "className":"Artificer",
      "classNameRU":"Изобретатель",
      "background":[
         "Guild artisan",
         "Sage",
         "Folk hero"
      ]
   }
]`)
