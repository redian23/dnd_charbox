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

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("html/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "PreGeneraTOR",
		})
	})
	router.GET("/dices", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dices.html", gin.H{
			"title": "PreGeneraTOR",
		})
	})
	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", gin.H{
			"content": "This is an about page...",
		})
	})

	router.RunTLS(":443", "./cert/certificate.pem", "./cert/key.pem")

}
