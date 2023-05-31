package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"pregen/api"
	"strings"
)

var htmlSitePath, htmlSiteRootPath string
var assetsSitePath, assetsSiteRootPath string

func init() {
	InitServerPathVars(true)
}

func InitServerPathVars(status bool) {
	if status == true {
		assetsSitePath = "/usr/share/nginx/html/assets"
		assetsSiteRootPath = "./usr/share/nginx/html/assets"
		htmlSitePath = "/usr/share/nginx/html/*.html"
		htmlSiteRootPath = "/usr/share/nginx/html/"
	} else {
		assetsSitePath = "web/assets"
		assetsSiteRootPath = "./web/assets"
		htmlSitePath = "web/html/*.html"
	}
}

func main() {
	router := gin.Default()
	// api method
	v1 := router.Group("api/v1/")
	{
		//v1.GET("/dice:number", func(c *gin.Context) {
		//	number := c.Param("number")
		//	api.GetDice(c, number)
		//})
		v1.GET("/roll", func(c *gin.Context) {
			api.GetMultiRoll(c)
		})
	}

	//site pages
	router.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	router.Static(assetsSitePath, assetsSiteRootPath)
	fmt.Println(htmlSitePath)
	router.LoadHTMLGlob(htmlSitePath)

	//router.GET("/", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", gin.H{
	//		"title": "PreGeneraTOR",
	//	})
	//})
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, htmlSiteRootPath+"dices.html", gin.H{
			"title": "PreGeneraTOR",
		})
	})
	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, htmlSiteRootPath+"about.html", gin.H{
			"content": "This is an about page...",
		})
	})
	router.GET("/version", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte("Dice don't exist."+"\n"))
	})

	router.RunTLS(":444", "/etc/letsencrypt/live/diceroll.swn.by/fullchain.pem",
		"/etc/letsencrypt/live/diceroll.swn.by/privkey.pem")
}
