package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"net/http"
	"os"
	"pregen/api"
	"strings"
)

const Version = "0.7.5 Beta build <DICE_Roll>"

func main() {
	InitServerPathVars(true)
	f, _ := os.Create(logPath + "diceroll_gin_errors.log")
	gin.DefaultErrorWriter = io.MultiWriter(f)

	router := gin.Default()

	// api method
	v1 := router.Group("api/v1/")
	{
		v1.GET("/roll", func(c *gin.Context) {
			api.GetMultiRoll(c)
		})
	}

	//site pages
	router.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})

	router.LoadHTMLGlob(htmlSitePath)
	router.StaticFS("/f", http.Dir(filePath))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dices.html", gin.H{
			"title": "Dice Roller",
			"path":  "./f/diceroll",
		})
	})
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	router.GET("/about", func(c *gin.Context) {
		api.GetAbout(c)
	})
	router.GET("/version", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(Version+" VK_RED23"+"\n"))
	})
	//router.Run(":8080") //local
	router.RunTLS(":410", "/etc/letsencrypt/live/diceroll.swn.by/fullchain.pem", "/etc/letsencrypt/live/diceroll.swn.by/privkey.pem") //prod
}
