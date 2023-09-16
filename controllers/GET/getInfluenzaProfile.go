package get

import (
	"net/http"

	"github.com/gin-gonic/gin"
	responseStruct "github.com/shivamsouravjha/influenza/structure/response"
)

func GetInfluenzaProfile(c *gin.Context) {
	username := c.Query("username")
	resp := responseStruct.SuccessResponse{}
	resp.Message = username
	resp.Status = "true"
	c.JSON(http.StatusOK, resp)
}
