package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/shivamsouravjha/influenza/config"
	"github.com/shivamsouravjha/influenza/constants"
	"github.com/shivamsouravjha/influenza/helpers/token"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		userToken := c.GetHeader("userToken")
		if userToken != "" {
			jwtCheck, err := token.VerifyToken(userToken, config.Get().JWT)
			if err != nil {
				c.AbortWithStatusJSON(401, constants.INVALID_TOKEN_RESPONSE)
				return
			}
			c.Set("email", jwtCheck.Value)
		} else {
			c.AbortWithStatusJSON(401, constants.INVALID_TOKEN_RESPONSE)
		}
	}
}
