package main

import (
	"github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"net/http"
	"os"
	"pregen/api"
	"pregen/pkg/db"
	"strings"
	"time"
)

const Version = "0.7.7 Beta build"

func main() {
	InitServerPathVars(true)
	db.PingMongoDB()

	f, _ := os.Create(logPath + "charbox_gin_errors.log")
	gin.DefaultErrorWriter = io.MultiWriter(f)

	router := gin.Default()

	// This makes it so each ip can only make 10 requests per second
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second,
		Limit: 10,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	// api method
	v1 := router.Group("api/v1/")
	{
		v1.GET("/get-character", mw, func(c *gin.Context) {
			api.GetRandomCharacter(c)
		})
		v1.POST("/post-current-character", mw, func(c *gin.Context) {
			api.GetCurrentCharacter(c)
		})
		v1.POST("/run-lss-export", mw, func(c *gin.Context) {
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
	//router.Run(":820") //local
	router.RunTLS(":420", "/etc/letsencrypt/live/charbox.swn.by/fullchain.pem", "/etc/letsencrypt/live/charbox.swn.by/privkey.pem") //prod
}

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
	time.Sleep(2 * time.Second)
}
