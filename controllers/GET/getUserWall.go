package get

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shivamsouravjha/influenza/helpers/mongo"
	responseStruct "github.com/shivamsouravjha/influenza/structure/response"
	"github.com/shivamsouravjha/influenza/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetInfluenzaWall(c *gin.Context) {
	userId := c.Query("userId")
	if userId == "" {
		c.JSON(422, utils.SendErrorResponse(errors.New("Can't find feedback")))
	}

	resp := responseStruct.SuccessResponse{}
	objID, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.M{"_id": objID}

	existingUser, err := mongo.GroupFind(filter, "inlfuenza")
	if err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return

	}
	resp.Message = existingUser
	resp.Status = "true"
	c.JSON(http.StatusOK, resp)
}
