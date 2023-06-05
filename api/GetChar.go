package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	cl "pregen/backend/classes"
)

type classApi struct {
	ClassName string `form:"class_name"`
}

func GetRandomClass(c *gin.Context) {
	c.JSONP(http.StatusOK, cl.GetNewClass())
}
