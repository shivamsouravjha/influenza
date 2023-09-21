package get

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shivamsouravjha/influenza/helpers/mongo"
	responseStruct "github.com/shivamsouravjha/influenza/structure/response"
	"github.com/shivamsouravjha/influenza/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func GetInfluenzaProfile(c *gin.Context) {
	username := c.Query("username")
	resp := responseStruct.SuccessResponse{}
	filter := bson.M{"_id": username}

	existingUser, err := mongo.Find(filter, "inlfuenza")
	if err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return

	}
	resp.Message = existingUser

	resp.Status = "true"
	c.JSON(http.StatusOK, resp)
}
