package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mazen160/go-random"
	"net/http"
	"pregen/pkg/characterCore"
)

type charParameters struct {
	ClassName string `json:"class"`
	RaceName  string `json:"race"`
}

func GetCurrentCharacter(c *gin.Context) {
	var chr charParameters
	if err := c.ShouldBindJSON(&chr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if chr.ClassName == "random" {
		classArray := []string{"Волшебник", "Алхимик", "Воин", "Варвар", "Паладин", "Монах",
			"Плут", "Следопыт", "Друид", "Жрец", "Чернокнижник", "Чародей", "Бард", "Изобретатель"}
		randNum, _ := random.IntRange(0, len(classArray)-1)
		chr.ClassName = classArray[randNum]
	}
	if chr.RaceName == "random" {
		raceArray := []string{"Аасимар", "Багбир", "Гном", "Гоблин", "Голиаф", "Дварф", "Кенку", "Кобольд",
			"Локата", "Орк", "Тифлинг", "Эльф", "Полуорк", "Полурослик", "Полуэльф", "Табакси", "Человек",
			"Мурлок", "Фирболг", "Драконорожденный", "Юань-ти", "Тритон"}
		randNum, _ := random.IntRange(0, len(raceArray)-1)
		chr.RaceName = raceArray[randNum]
	}

	c.JSONP(http.StatusOK, characterCore.GenerateFullCharacter(chr.ClassName, chr.RaceName))

}
