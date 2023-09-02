package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shivamsouravjha/influenza/services"
	requestStruct "github.com/shivamsouravjha/influenza/structure/request"
	responseStruct "github.com/shivamsouravjha/influenza/structure/response"
	"github.com/shivamsouravjha/influenza/utils"
)

func VerifyCode(c *gin.Context) {
	verifyRequest := requestStruct.OTPRequest{}
	if err := c.ShouldBind(&verifyRequest); err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}

	err := services.Verifycode(verifyRequest)

	if err != nil {
		resp := responseStruct.SuccessResponse{}
		resp.Message = err.Error()
		resp.Status = "false"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := responseStruct.SuccessResponse{}
	resp.Message = "Email verified successfully"
	resp.Status = "true"
	c.JSON(http.StatusOK, resp)
}
