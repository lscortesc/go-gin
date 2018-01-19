package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/lscortesc/go-gin/controllers"
	"gitlab.com/lscortesc/go-gin/db"
	"gitlab.com/lscortesc/go-gin/models"
)

func setupRouter() *gin.Engine {

	r := gin.Default()

	// Default
	defaultCtrl := new(controllers.DefaultController)
	r.GET("/", defaultCtrl.HomeAction)

	// Auth
	authCtrl := new(controllers.AuthController)
	r.POST("/register", authCtrl.Register)

	return r
}

func main() {

	db.Setup()
	makeMigrations()

	r := setupRouter()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}

func makeMigrations() {
	db.GetConnection().AutoMigrate(
		&models.Country{},
		&models.Customer{},
	)

	models.Customer{}.MigrateRelationships(db.GetConnection())
	models.Country{}.Seed(db.GetConnection())
}
