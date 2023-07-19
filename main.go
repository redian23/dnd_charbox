package main

import (
	"github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"net/http"
	"os"
	"pregen/api"
	"pregen/db"
	"strings"
	"time"
)

var htmlSitePath string
var imagesSitePath = "/images"
var imagesSiteRootPath string
var logPath string

func init() {
	InitServerPathVars(true)
	//InitServerPathVars(false)
	db.PingMongoDB()
}

func InitServerPathVars(status bool) {
	if status == true {
		imagesSiteRootPath = "./usr/share/nginx/html/images"
		htmlSitePath = "/usr/share/nginx/html/*.html"
		logPath = "/var/logs/"
	} else {
		imagesSiteRootPath = "./frontend/images"
		htmlSitePath = "frontend/html/*.html"
		logPath = ""
	}
}

func main() {
	f, _ := os.Create(logPath + "gin_errors.log")
	gin.DefaultErrorWriter = io.MultiWriter(f)

	router := gin.Default()

	// This makes it so each ip can only make 5 requests per second
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second,
		Limit: 5,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	// api method
	v1 := router.Group("api/v1/")
	{
		//v1.GET("/dice:number", func(c *gin.Context) {
		//	number := c.Param("number")
		//	api.GetDice(c, number)
		//})
		v1.GET("/roll", mw, func(c *gin.Context) {
			api.GetMultiRoll(c)
		})

		v1.GET("/get-character", mw, func(c *gin.Context) {
			api.GetRandomCharacter(c)
		})

	}

	//site pages
	router.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	router.Static(imagesSitePath, imagesSiteRootPath)
	router.LoadHTMLGlob(htmlSitePath)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dices.html", gin.H{
			"title": "Dice Roller",
		})
	})
	router.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "character_box_ru.html", gin.H{
			"title": "Character Box",
		})
	})
	router.GET("/about", func(c *gin.Context) {
		api.GetAbout(c)
	})
	router.GET("/version", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(Version+" VK_RED23"+"\n"))
	})
	//router.Run(":848") //local
	router.RunTLS(":444", "/etc/letsencrypt/live/diceroll.swn.by/fullchain.pem", "/etc/letsencrypt/live/diceroll.swn.by/privkey.pem") //prod
}

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
	time.Sleep(2 * time.Second)
}
