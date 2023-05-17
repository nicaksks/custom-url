package routes

import (
	"net/http"
	"short-url/backend/database"
	"short-url/backend/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	domain = utils.Domain()
)

func Register(c *gin.Context) {
	_, err := database.Find(c.PostForm("short"))
	if err == mongo.ErrNoDocuments {
		err := database.Register(c.PostForm("short"), c.PostForm("url"))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": "Error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Success!",
			"info": gin.H{
				"domain":   domain,
				"short":    c.PostForm("short"),
				"url":      domain + c.PostForm("short"),
				"redirect": c.PostForm("url"),
			},
		})
		return
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusOK,
			"message": "This name is already registered.",
		})
		return
	}
}
