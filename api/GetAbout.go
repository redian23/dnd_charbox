package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type about struct {
	Projects []string          `json:"projects"`
	Author   string            `json:"author"`
	Team     map[string]string `json:"team"`
	Thanks   map[string]string `json:"thanks"`
}

func GetAbout(c *gin.Context) {
	aboutCharBox := about{
		Projects: []string{"Character Box", "DiceRoll"},
		Author:   "Redian23",
		Team: map[string]string{
			"Admin/FrontEnd/BackEnd": "Redian23",
			"QA":                     "Racist Ooga Booga",
			"A_QA":                   "Lisha_Svitok",
		},
		Thanks: map[string]string{
			"Special thanks to": "shakusky.lss",
		},
	}

	c.JSONP(http.StatusOK, aboutCharBox)
}
