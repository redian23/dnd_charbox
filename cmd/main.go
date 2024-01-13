package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"pregen/api"
	"pregen/pkg/db"
	"strings"
)

const Version = "0.9.4.3 Beta build"

func main() {
	ProdStatus = true
	InitServerPathVars()
	db.PingMongoDB()

	router := gin.Default()
	// api method
	v1 := router.Group("api/v1/")
	{
		v1.POST("/post-current-character", func(c *gin.Context) {
			api.GetCurrentCharacter(c)
		})
		v1.POST("/run-lss-export", func(c *gin.Context) {
			api.GetLSSJson(c)
		})
	}

	//site pages
	router.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})

	router.LoadHTMLGlob(htmlSitePath)
	router.StaticFile("favicon.png", filePath+"/charbox/favicon.png")
	router.StaticFile("charbox.js", filePath+"/charbox/charbox.js")
	router.StaticFile("charbox.css", filePath+"/charbox/charbox.css")
	router.StaticFS("/imgs", http.Dir(imgsPath))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "character_box_ru.html", gin.H{
			"title":   "Шкатулка Персонажей | Character Box | Генератор персонажей для DnD 5e",
			"version": Version,
		})
	})

	router.GET("/faq", func(c *gin.Context) {
		c.HTML(http.StatusOK, "charbox_faq.html", gin.H{
			"title":   "Шкатулку Персонажей - Вопросы/Ответы | Character Box - FAQ",
			"version": Version,
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})
	router.GET("/about", func(c *gin.Context) {
		api.GetAbout(c)
	})
	router.GET("/s_map.xml", func(c *gin.Context) {
		c.File(filePath + "/charbox/s_map.xml")
	})
	router.GET("/robots.txt", func(c *gin.Context) {
		c.File(filePath + "/charbox/robots.txt")
	})

	if ProdStatus == true {
		router.RunTLS(":420",
			"/etc/letsencrypt/live/charbox.swn.by/fullchain.pem",
			"/etc/letsencrypt/live/charbox.swn.by/privkey.pem") //prod
	} else {
		router.Run(":8080") //local
	}
}
