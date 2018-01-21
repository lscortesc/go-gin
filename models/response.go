package models

import (
	"strings"

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

// GetErrors from validation (errMsgs)
func (r *APIResponse) GetErrors(c *gin.Context) map[string]interface{} {
	var err []string

	if c.Errors.Errors()[0] == "EOF" {
		err = append(err, "No Data Provided")
		return gin.H{"errors": err}
	}

	errors := strings.Split(c.Errors.Errors()[0], "\n")

	for _, e := range errors {
		err = append(err, strings.Split(e, "' ")[1])
	}

	return gin.H{"errors": err}
}
