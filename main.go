package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/lscortesc/go-gin/db"
	"gitlab.com/lscortesc/go-gin/models"
)

func setupRouter() *gin.Engine {

	db.Setup()

	r := gin.Default()

	// test
	r.GET("/", func(c *gin.Context) {
		res := models.APIResponse{
			Status:  http.StatusOK,
			Message: "Data recover successfully",
		}

		res.Response(c)
	})

	return r
}

func main() {
	r := setupRouter()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
