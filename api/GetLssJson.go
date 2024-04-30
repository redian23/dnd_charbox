package api

//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"pregen/pkg/characterCore"
//	"pregen/pkg/export"
//)
//
//func GetLSSJson(c *gin.Context) {
//	var char characterCore.Character
//	if err := c.ShouldBindJSON(&char); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSONP(http.StatusOK, export.RunExportToLSS(char))
//}
