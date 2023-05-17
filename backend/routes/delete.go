package routes

import (
	"net/http"
	"short-url/backend/database"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	err := database.Delete(c.Param("short"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Successfully deleted!",
	})
	return
}
