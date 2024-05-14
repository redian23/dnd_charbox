package backgrounds

import (
	"encoding/json"
	"log"
	"os"
	"pregen/pkg/db"
)

var backgroundName string

func getBackgroundsFormDB() []BackgroundBson {
	var results []BackgroundBson
	fileContent, err := os.Open(db.DataBasePath + "backgrounds.json")
	if err != nil {
		log.Println(err)
	}
	defer fileContent.Close()
	var byteResult, _ = os.ReadFile(db.DataBasePath + "backgrounds.json")

	json.Unmarshal(byteResult, &results)
	return results
}
