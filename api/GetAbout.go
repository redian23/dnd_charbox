package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type about struct {
	Projects []string          `json:"projects"`
	Author   string            `json:"author"`
	Thank    map[string]string `json:"thank"`
}

func GetAbout(c *gin.Context) {
	aboutCharBox := about{
		Projects: []string{"Character Box", "DiceRoll"},
		Author:   "Redian23",
		Thank: map[string]string{
			"QA": "Racist Ooga Booga",
		},
	}

	c.JSONP(http.StatusOK, aboutCharBox)
}
