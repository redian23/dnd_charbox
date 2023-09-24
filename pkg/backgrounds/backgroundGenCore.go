package backgrounds

import (
	"context"
	"github.com/mazen160/go-random"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"pregen/pkg/classes"
	"pregen/pkg/db"
)

var backgroundName string

func getBackgroundsFormDB() []BackgroundBson {
	var results []BackgroundBson

	client := db.ConnectToDB()
	ctx := context.Background()

	coll := client.Database(db.DBNAME).Collection("backgrounds")
	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	defer func(client *mongo.Client, ctx context.Context) {
		err = client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(client, ctx)

	return results
}

func backgroundAnalyze(classNameRu string) string {
	var chars = classes.GetClassCharactsFormDB()
	for _, char := range chars {
		if classNameRu == char.ClassNameRU {
			var rollNum int
			rollNum, _ = random.IntRange(0, len(char.Background))

			backgroundName = char.Background[rollNum] //get Background name for all methods in package
			return char.Background[rollNum]
		}
	}
	return ""
}
