package backgrounds

import (
	"context"
	"github.com/mazen160/go-random"
	"pregen/pkg/classes"
	"pregen/pkg/db"
)

var backgroundName string

func getBackgroundsFormDB() []BackgroundBson {
	var results []BackgroundBson
	var cursor = db.ReadFromDB("backgrounds")

	if err := cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return results
}

func backgroundAnalyze(className string) string {
	var chars = classes.GetClassCharactsFormDB()
	for _, char := range chars {
		if className == char.ClassName {
			var rollNum int
			rollNum, _ = random.IntRange(0, len(char.Background))

			backgroundName = char.Background[rollNum] //get Background name for all methods in package
			return char.Background[rollNum]
		}
	}
	return ""
}
