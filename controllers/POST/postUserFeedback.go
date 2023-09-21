package post

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	cloudinary "github.com/shivamsouravjha/influenza/helpers/cloudinary"
	"github.com/shivamsouravjha/influenza/helpers/mongo"
	"github.com/shivamsouravjha/influenza/structure"
	requestStruct "github.com/shivamsouravjha/influenza/structure/request"
	responseStruct "github.com/shivamsouravjha/influenza/structure/response"
	"github.com/shivamsouravjha/influenza/utils"
	"go.mongodb.org/mongo-driver/bson"
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
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error:", err)
			c.JSON(422, utils.SendErrorResponse(err))
			return
		}
		imagePath := filepath.Join(currentDir, image)
		uploaded, err := cloudinary.Client().Upload.Upload(context.Background(), imagePath, uploader.UploadParams{PublicID: image})
		if err != nil {
			fmt.Println("Error", err.Error())
			c.JSON(422, utils.SendErrorResponse(err))
			return
		}
		err = os.Remove(imagePath)
		if err != nil {
			fmt.Println("Error deleting file:", err)
			c.JSON(422, utils.SendErrorResponse(err))
			return
		}
		influencerFeedback.ImageLink = uploaded.SecureURL
		influencerFeedback.Post = nil
	}

	resp := responseStruct.SuccessResponse{}

	// email, _ := c.Get("email")

	// if feedBack, ok := email.(string); ok {
	// 	resp.Message = feedBack
	influenzaInfo := structure.Feedbackdata{
		Name:     influencerFeedback.Name,
		Company:  influencerFeedback.Company,
		LinkedIn: influencerFeedback.LinkedIn,
	}

	filter := bson.M{"linkedin": influencerFeedback.LinkedIn}

	existingUser, err := mongo.FindInfluenza(filter, "inlfuenza")
	if err != nil {
		res, createError := mongo.CreateInfluenza(influenzaInfo, "inlfuenza")
		if createError != nil {
			c.JSON(422, utils.SendErrorResponse(err))
			return
		}
		influencerFeedback.Influencer = res.InsertedID
		createError = mongo.CreateItem(influencerFeedback, "feedback")
		if createError != nil {
			c.JSON(422, utils.SendErrorResponse(err))
			return
		}
	} else {
		influencerFeedback.Influencer = &existingUser.ID
		createError := mongo.CreateItem(influencerFeedback, "feedback")
		if createError != nil {
			c.JSON(422, utils.SendErrorResponse(err))
			return
		}
	}
	resp.Status = "succesful"
	resp.Message = "feedback submitted successfully"
	c.JSON(http.StatusOK, resp)
}
