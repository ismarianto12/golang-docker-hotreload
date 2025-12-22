package utils

import "github.com/gin-gonic/gin"

type UploadFile struct {
	File     interface{}
	Filename string
}

func ActionUplooadFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// fileUpload := c.
	}
}
