package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"pregen/api"
	"strings"
)

var port string
var titleText string
var siteMap string

func setStage(stage string) {
	switch stage {
	case "prod":
		port = ":4050"
		titleText = "Шкатулка Персонажей | Character Box | Генератор персонажей для DnD 5e"
		siteMap = "s_map.xml"
	case "test":
		port = ":4090"
		titleText = "Шкатулка Персонажей Тест | Character Box Beta"
		siteMap = "s_map_test.xml"
	}
}

func main() {
	setStage("test")

	router := gin.Default()
	// api method
	v1 := router.Group("api/v2/")
	{
		v1.POST("/get-character", func(c *gin.Context) {
			api.GetCurrentCharacter(c)
		})
	}
	//site pages
	router.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})

	router.LoadHTMLGlob(GlobPattern)
	router.StaticFile("favicon.png", WebPagesPath+"/favicon.png")
	router.StaticFile("charbox.js", WebPagesPath+"/charbox.js")
	router.StaticFile("charbox-pico.css", WebPagesPath+"/charbox-pico.css")
	router.StaticFS("/img", http.Dir(ImagePath))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "character_box_pico.html", gin.H{
			"title":   titleText,
			"version": Version,
		})
	})

	router.GET("/faq", func(c *gin.Context) {
		c.HTML(http.StatusOK, "charbox_faq.html", gin.H{
			"title": "Раздел вопросов и ответов | Q&A Page",
		})
	})

	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "charbox_about.html", gin.H{
			"title": "Наша команда | Our Team",
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	router.GET("/s_map.xml", func(c *gin.Context) {
		c.File(WebPagesPath + siteMap)
	})
	router.GET("/robots.txt", func(c *gin.Context) {
		c.File(WebPagesPath + "robots.txt")
	})

	router.Run(port)
}
