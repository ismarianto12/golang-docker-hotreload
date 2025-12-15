package handlers

import (
	"log"

	"encoding/base64"

	"github.com/gin-gonic/gin"
)

type HeaderReturn struct {
	Response bool
}

func DecodeToken(param string) string {
	result, err := base64.StdEncoding.DecodeString(param)
	if err != nil {
		return err.Error()
	}
	decoded := string(result)
	return decoded

}

func CheckTokenHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		authheader := c.GetHeader("Authorization")
		param := DecodeToken(authheader)
		log.Print("decode token result %s", param)

		if authheader == "" {
			c.JSON(403, gin.H{
				"ms": "Authorized",
			})
			c.Abort()
			return
		}
		// token := strings.Split(authheader, "")
		log.Print("detil %s", param)
		c.Next()
	}

}
