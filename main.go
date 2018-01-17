package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/lscortesc/go-gin/models"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
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
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
