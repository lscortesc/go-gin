package controllers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/lscortesc/go-gin/forms"
	"gitlab.com/lscortesc/go-gin/models"
)

// AuthController Authentication Controller
type AuthController struct{}

var customerModel models.Customer
var response models.APIResponse
var err error

// Register Action
func (ctrl AuthController) Register(c *gin.Context) {
	var registerForm forms.RegisterForm
	var customer models.Customer

	if c.BindJSON(&registerForm) != nil {
		response.Response(c, c.Errors, 422, "Data Not Valid")
		return
	}

	if customer, err = customerModel.Register(registerForm); err != nil {
		response.Response(c, gin.H{"error": err.Error()}, 400, "Error at register customer")
		return
	}

	response.Response(c, customer, 201, "Customer Registerd")
}
