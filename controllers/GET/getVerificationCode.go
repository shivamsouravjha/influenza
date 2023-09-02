package get

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shivamsouravjha/influenza/constants"
	"github.com/shivamsouravjha/influenza/helpers/email"
	responseStruct "github.com/shivamsouravjha/influenza/structure/response"
)

func VerificationCode(c *gin.Context) {
	requestEmail := c.Query("email")
	err := email.SendEmail(requestEmail, constants.VerifyYourself)

	if err != nil {
		resp := responseStruct.SuccessResponse{}
		resp.Message = err.Error()
		resp.Status = "false"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := responseStruct.SuccessResponse{}
	resp.Message = "OTP sent successfully"
	resp.Status = "true"
	c.JSON(http.StatusOK, resp)
}
