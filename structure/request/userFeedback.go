package requestStruct

import "mime/multipart"

type InfluencerFeedback struct {
	Name     string                `form:"name" binding:"required"`
	LinkedIn string                `form:"linkedin" binding:"required"`
	Post     *multipart.FileHeader `form:"post" binding:"required"`
	Company  string                `form:"company" binding:"required"`
}
