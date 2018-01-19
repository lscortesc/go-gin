package models

import (
	"github.com/gin-gonic/gin"
)

// APIResponse returns an response structured
type APIResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
}

// Response generic
func (r *APIResponse) Response(c *gin.Context, response interface{}, status int, message string) {

	r.Data = response
	r.Status = status
	r.Message = message

	c.JSON(r.Status, r)
}
