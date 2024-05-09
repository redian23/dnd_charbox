package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"pregen/pkg/core"
	"strings"
)

func main() {
	core.GetFullCharacter()

	router := gin.Default()
	// api method
	v1 := router.Group("api/v1/")
	{
		v1.POST("/post-current-character", func(c *gin.Context) {
			//api.GetCurrentCharacter(c)
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
	router.StaticFS("/imgs", http.Dir(ImagePath))

	router.GET("/preview", func(c *gin.Context) {
		c.HTML(http.StatusOK, "character_box_pico.html", gin.H{
			"title":   "Шкатулка Персонажей | Character Box | Генератор персонажей для DnD 5e",
			"version": Version,
		})
	})
	router.Run(":4090")
}
