package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type about struct {
	Projects []string          `json:"projects"`
	Author   string            `json:"author"`
	Team     []team            `json:"team"`
	Thanks   map[string]string `json:"thanks"`
}
type team struct {
	Occupation string `json:"occupation"`
	Name       string `json:"name"`
}

func GetAbout(c *gin.Context) {
	aboutCharBox := about{
		Projects: []string{"Character Box", "DiceRoll"},
		Author:   "Redian23",
		Team: []team{
			{Occupation: "SysAdmin/BackEnd/FrontEnd", Name: "Redian23"},
			{Occupation: "QA/PR Manager", Name: "Racist Ooga Booga"},
			{Occupation: "A_QA", Name: "Lisha_Svitok"},
		},
		Thanks: map[string]string{
			"Special thanks to": "shakusky.lss",
		},
	}

	c.JSONP(http.StatusOK, aboutCharBox)
}
