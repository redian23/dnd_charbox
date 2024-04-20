package races

import (
	"fmt"
	"github.com/mazen160/go-random"
)

var RaceInfo *Race

func GetRaceFirstName(gender string, raceNames map[string][]string) string {
	randNum, _ := random.IntRange(0, len(raceNames[gender]))
	return raceNames[gender][randNum]
}

func GetRaceLastName(raceLastNames []string) string {
	randNum, _ := random.IntRange(0, len(raceLastNames))
	return raceLastNames[randNum]
}

func GetBodyStats(stats BodyStats) *BodyAnswer {
	var eyesColors = []string{"Желтые", "Серые", "Черные", "Синие", "Красные", "Зеленые", "Карие", "Морская волна"}
	eyesColor, _ := random.Choice(eyesColors)

	var hairColors = []string{"Белые", "Черные", "Синие", "Красные", "Зеленые", "Каштановые", "Бирюзовые", "Рыжие"}
	hairColor, _ := random.Choice(hairColors)

	raceAge, _ := random.IntRange(stats.MinAge, (stats.MaxAge*75)/100)
	raceHeight, _ := random.IntRange(stats.MinHeight, stats.MaxHeight)
	raceWeight, _ := random.IntRange(stats.MinWeight, stats.MaxWeight)

	return &BodyAnswer{
		Age:      raceAge,
		Height:   raceHeight,
		Weight:   raceWeight,
		HeightFt: fmt.Sprintf("%.1f", float32(raceHeight)*0.0328),
		WeightLb: (raceWeight * 220) / 100,
		BodySize: stats.BodySize,
		Eyes:     eyesColor,
		Hair:     hairColor,
	}
}
