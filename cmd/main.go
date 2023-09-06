package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"pregen/api"
	"pregen/pkg/db"
	"strings"
)

const Version = "0.8.4 Beta build"

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
	router.StaticFS("/f", http.Dir(filePath))
	router.StaticFS("/imgs", http.Dir(imgsPath))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "character_box_ru.html", gin.H{
			"title": "Character Box",
			"path":  "./f/charbox",
		})
	})

	router.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "character_box_test.html", gin.H{
			"title": "CharBox Beta",
			"path":  "./f/charbox",
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})
	router.GET("/about", func(c *gin.Context) {
		api.GetAbout(c)
	})

	router.GET("/version", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(Version+" VK_RED23"+"\n"))
	})

	if ProdStatus == true {
		router.RunTLS(":420", "/etc/letsencrypt/live/charbox.swn.by/fullchain.pem", "/etc/letsencrypt/live/charbox.swn.by/privkey.pem") //prod
	} else {
		router.Run(":8080") //local
	}
}
