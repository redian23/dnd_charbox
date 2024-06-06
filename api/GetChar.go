package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mazen160/go-random"
	"net/http"
	"pregen/pkg/core"
)

type charParameters struct {
	ClassName          string `json:"class"`
	ClassArchetypeName string `json:"class_archetype"`
	RaceName           string `json:"race"`
	RaceTypeName       string `json:"race_type"`
	BackgroundName     string `json:"background"`
	Level              int    `json:"level"`
	Gender             string `json:"gender"`
}

var classArray = []string{"Варвар", "Бард"}

var raceArray = []string{"Аасимар", "Багбир", "Гном", "Темный эльф", "Гоблин", "Голиаф", "Дворф",
	"Кенку"}

var backgroundArray = []string{
	"Артист", "Беспризорник", "Благородный",
	"Гильдейский ремесленник", "Моряк",
	"Мудрец", "Пират", "Рыцарь",
}

func GetCurrentCharacter(c *gin.Context) {
	var chr charParameters
	if err := c.ShouldBindJSON(&chr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if chr.ClassName == "random" {
		randNum, _ := random.IntRange(0, len(classArray))
		chr.ClassName = classArray[randNum]
	}
	if chr.ClassArchetypeName == "random" || chr.ClassArchetypeName == "" {
		chr.ClassArchetypeName = "Стандартный"
	}

	if chr.RaceName == "random" {
		randNum, _ := random.IntRange(0, len(raceArray))
		chr.RaceName = raceArray[randNum]
	}

	//if chr.RaceTypeName == "random" || chr.RaceTypeName == "" {
	//	chr.RaceTypeName = "Стандартный"
	//}

	if chr.BackgroundName == "random" {
		randNum, _ := random.IntRange(0, len(backgroundArray))
		chr.BackgroundName = backgroundArray[randNum]
	}
	if chr.Level > 8 {
		chr.Level = 8
	}

	c.JSONP(http.StatusOK, core.GetFullCharacter(
		chr.ClassName,
		chr.ClassArchetypeName,
		chr.RaceName,
		chr.RaceTypeName,
		chr.BackgroundName,
		chr.Gender,
		chr.Level,
	))
}
