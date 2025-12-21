package middlewares

import (
	"os"
	"rianRestapp/entities"
	"rianRestapp/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type ParsingJwt struct {
	secretkey string
}

func NewJwtSigned() *ParsingJwt {
	return &ParsingJwt{
		secretkey: os.Getenv("SECRET_KEY"),
	}
}

func (md *ParsingJwt) GenerateToken(UserId uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": UserId,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES384, claims)
	return token.SignedString(md.secretkey)
}

func (md *ParsingJwt) AuthToken(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authtoken := c.GetHeader("Authorization")
		if authtoken == "" {
			// c.JSON()
			utils.BuildResponse(nil, 402, "Unatithorize", c)
			c.Abort()
			return
		}

		// chehk kombinasi bearer token
		bearertoken := strings.Split(authtoken, " ")
		if len(bearertoken) != 2 || bearertoken[0] != "Bearer" {
			utils.BuildResponse(nil, 402, "Invalid Token Access", c)
			c.Abort()
			return

		}
		tokenStr := bearertoken[1]

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return md.secretkey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := uint(claims["user_id"].(float64))

		// 🔥 VALIDASI KE TABLE USERS
		var user entities.User
		if err := db.First(&user, userID).Error; err != nil {
			c.JSON(401, gin.H{"message": "User not active"})
			c.Abort()
			return
		}

		c.Set("access_token", tokenStr)
		c.Next()
	}

}
