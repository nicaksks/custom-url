package backend

import (
	"fmt"
	"log"
	"net/http"
	"short-url/backend/routes"
	"short-url/backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	port = utils.Port()
)

func Router() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.LoadHTMLGlob("./frontend/pages/*.html")
	router.Static("/static", "./frontend")

	router.GET(":short", routes.Index)
	router.GET("/create", routes.Create)
	router.POST("/register", routes.Register)
	router.DELETE("/:short", routes.Delete)

	return router
}

func Start() {
	server := &http.Server{
		Addr:    port,
		Handler: Router(),
	}

	fmt.Printf("Server online! using port %s\n", strings.Replace(port, ":", "", 1))
	if err := server.ListenAndServe(); err != nil {
		log.Printf("Error: %s", err)
	}
}
