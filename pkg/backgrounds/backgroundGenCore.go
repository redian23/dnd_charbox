package backgrounds

import (
	"encoding/json"
	"github.com/mazen160/go-random"
	"log"
	"os"
	"pregen/pkg/classes"
)

var backgroundName string

func getBackgroundsFormDB() []BackgroundBson {
	var results []BackgroundBson
	fileContent, err := os.Open("/etc/charbox.d/db/backgrounds.json")
	if err != nil {
		log.Println(err)
	}
	defer fileContent.Close()
	var byteResult, _ = os.ReadFile("/etc/charbox.d/db/backgrounds.json")

	json.Unmarshal(byteResult, &results)
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
