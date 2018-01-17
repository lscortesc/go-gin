package models

import (
	"github.com/gin-gonic/gin"
)

// APIResponse returns an response structured
type APIResponse struct {
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
	Status  int                    `json:"status"`
}

// Response generic
func (r *APIResponse) Response(c *gin.Context) {
	if len(r.Data) == 0 {
		r.Data = map[string]interface{}{}
	}

	c.JSON(r.Status, r)
}
