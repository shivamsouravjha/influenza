package post

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	requestStruct "github.com/shivamsouravjha/influenza/structure/request"
	responseStruct "github.com/shivamsouravjha/influenza/structure/response"
	"github.com/shivamsouravjha/influenza/utils"
)

func PostFeedback(c *gin.Context) {
	influencerFeedback := requestStruct.InfluencerFeedback{}
	if err := c.ShouldBind(&influencerFeedback); err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}

	if influencerFeedback.Post != nil {
		uniqueId := uuid.New()
		filename := strings.Replace(uniqueId.String(), "-", "", -1)
		fileExt := strings.Split(influencerFeedback.Post.Filename, ".")[1]
		image := fmt.Sprintf("%s.%s", filename, fileExt)
		err := c.SaveUploadedFile(influencerFeedback.Post, image)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	resp := responseStruct.SuccessResponse{}
	message, _ := c.Get("email")
	if intValue, ok := message.(string); ok {
		// Successfully converted to int, now convert to string
		resp.Message = intValue
		fmt.Printf("Integer: %d, String: %s\n", intValue)
	} else {
		fmt.Println("Value is not an int")
	}
	c.JSON(http.StatusOK, resp)

}
