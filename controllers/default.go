package controllers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/lscortesc/go-gin/models"
)

// DefaultController type
type DefaultController struct{}

// HomeAction Default Action
func (ctrl DefaultController) HomeAction(c *gin.Context) {
	res := new(models.APIResponse)

	res.Response(c, gin.H{}, 200, "Default Controller | Home Action")
}
