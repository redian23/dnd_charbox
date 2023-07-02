package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pregen/backend/characterCore"
)

func GetRandomCharacter(c *gin.Context) {
	c.JSONP(http.StatusOK, characterCore.StartCharacterGenerate())
}
