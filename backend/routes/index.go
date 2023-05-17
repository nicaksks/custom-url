package routes

import (
	"net/http"
	"short-url/backend/database"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	find, err := database.Find(c.Param("short"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusOK,
			"message": "This address does not exist.",
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, find.Url)
}