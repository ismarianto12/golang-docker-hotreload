package utils

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	data    interface{} `json:"data,omitempty"`
	code    int         `json:"code"`
	message string      `json:"message"`
}

func BuildResponse(data interface{}, code int, message string, c *gin.Context) {

	c.JSON(code, gin.H{
		"data":    data,
		"code":    code,
		"message": message,
		"time":    time.Now(),
	})
	return
}

const SUCCESS string = "SUCCESS FULLY RETREIVE DATA"
const ERRORRETRIVE string = "SUCCESS FULLY RETREIVE DATA"
