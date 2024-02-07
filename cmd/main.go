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

const Version = "1.0 Stable build <DICE_Roll>"

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

	router.GET("/ru", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dices_ru.html", gin.H{
			"title": "Dice Roller | Генератор броска кубика DnD 5e",
			"path":  "./f/diceroll",
		})
	})
	router.GET("/eng", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dices_eng.html", gin.H{
			"title": "Dice Roller | DnD 5e die roll generator",
			"path":  "./f/diceroll",
		})
	})
	router.GET("/s_map.xml", func(c *gin.Context) {
		c.File(filePath + "/diceroll/s_map.xml")
	})
	router.GET("/robots.txt", func(c *gin.Context) {
		c.File(filePath + "/diceroll/robots.txt")
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

	router.Run(":4040")
}
