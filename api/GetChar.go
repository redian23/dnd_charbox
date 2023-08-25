package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pregen/pkg/characterCore"
)

type charParameters struct {
	ClassName string `json:"class"`
	RaceName  string `json:"race"`
}

func GetRandomCharacter(c *gin.Context) {
	c.JSONP(http.StatusOK, characterCore.GenerateFullCharacter())
}

func GetCurrentCharacter(c *gin.Context) {
	var chr charParameters
	if err := c.ShouldBindJSON(&chr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for {
		var currentCharacter = characterCore.GenerateFullCharacter()

		if chr.ClassName == currentCharacter.Class.ClassNameRU &&
			chr.RaceName == currentCharacter.Race.RaceNameRu {
			c.JSONP(http.StatusOK, currentCharacter)
			break
		}
		if chr.ClassName == currentCharacter.Class.ClassNameRU && chr.RaceName == "random" ||
			chr.RaceName == currentCharacter.Race.RaceNameRu && chr.ClassName == "random" {
			c.JSONP(http.StatusOK, currentCharacter)
			break
		}
	}
}
