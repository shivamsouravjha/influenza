package delete

import (
	"errors"
	"net/http"

	responseStruct "github.com/shivamsouravjha/influenza/structure/response"

	"github.com/gin-gonic/gin"
	"github.com/shivamsouravjha/influenza/helpers/mongo"
	"github.com/shivamsouravjha/influenza/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeletePost(c *gin.Context) {

	feedbackId := c.Query("feedbackId")
	if feedbackId == "" {
		c.JSON(422, utils.SendErrorResponse(errors.New("Can't find feedback")))
	}

	resp := responseStruct.SuccessResponse{}
	objID, _ := primitive.ObjectIDFromHex(feedbackId)

	filter := bson.M{"_id": objID}

	existingUser, err := mongo.Delete(filter, "feedback")
	if err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return

	}
	resp.Message = existingUser
	resp.Status = "true"
	c.JSON(http.StatusOK, resp)
}
