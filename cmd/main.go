package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"pregen/api"
	"strings"
)

func main() {
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

	router.LoadHTMLGlob(GlobPattern)
	router.StaticFile("favicon.png", WebPagesPath+"/favicon.png")
	router.StaticFile("charbox.js", WebPagesPath+"/charbox.js")
	router.StaticFile("charbox.css", WebPagesPath+"/charbox.css")
	router.StaticFS("/imgs", http.Dir(ImagePath))

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
		c.File(WebPagesPath + "/s_map.xml")
	})
	router.GET("/robots.txt", func(c *gin.Context) {
		c.File(WebPagesPath + "/robots.txt")
	})

	router.Run(":4050")
}
