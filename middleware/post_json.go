package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidatePostJSON(c *gin.Context) {
	var jsonData interface{}
	x, _ := c.GetRawData()
	err := json.Unmarshal(x, &jsonData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"error":   error.Error(err),
			"message": "Bad JSON.",
		})
		c.Abort()
		return
	}
	c.Set("json", jsonData)
	c.Set("jsonByte", x)
}
