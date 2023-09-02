package get

import (
	"github.com/gin-gonic/gin"
	"github.com/shivamsouravjha/influenza/constants"
	"github.com/shivamsouravjha/influenza/helpers"
)

func VerificationCode(c *gin.Context) {
	email := c.Query("email")
	helpers.SendEmail(email, constants.VerifyYourself)
	c.JSON(200, "Worked verification")
}
