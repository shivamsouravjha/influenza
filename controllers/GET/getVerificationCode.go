package get

import (
	"github.com/gin-gonic/gin"
	"github.com/shivamsouravjha/influenza/constants"
	"github.com/shivamsouravjha/influenza/services"
)

func VerificationCode(c *gin.Context) {
	email := c.Query("email")
	services.SendEmail(email, constants.VerifyYourself)
	c.JSON(200, "Worked verification")
}
