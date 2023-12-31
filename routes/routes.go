package routes

import (
	"github.com/gin-gonic/gin"
	delete "github.com/shivamsouravjha/influenza/controllers/DELETE"
	get "github.com/shivamsouravjha/influenza/controllers/GET"
	post "github.com/shivamsouravjha/influenza/controllers/POST"
	"github.com/shivamsouravjha/influenza/helpers/cloudinary"
	"github.com/shivamsouravjha/influenza/helpers/mongo"
	"github.com/shivamsouravjha/influenza/helpers/redis"
	"github.com/shivamsouravjha/influenza/middlewares"
)

func NewRouter() *gin.Engine {
	redis.RedisInit()
	mongo.InitDb()
	cloudinary.Init()
	router := gin.New()
	v1 := router.Group("/api")

	emailRoutes := v1.Group("/email")
	{
		emailRoutes.GET("/getVerificationCode", get.VerificationCode)
		emailRoutes.POST("/verifyCode", post.VerifyCode)
	}
	postRoutes := v1.Group("/post")
	{
		postRoutes.GET("/getFeedback", get.GetUserFeedback)
		// postRoutes.Use(middlewares.JWT())
		postRoutes.POST("/postFeedback", post.PostFeedback)
	}
	userRoutes := v1.Group("/inluenza")
	{
		userRoutes.GET("/wall", get.GetInfluenzaWall)
		userRoutes.POST("/profile", get.GetInfluenzaProfile)
	}
	superAdmin := v1.Group("/superAdmin")
	{
		superAdmin.Use(middlewares.SuperJWT())
		superAdmin.DELETE("/post", delete.DeletePost)
	}
	return router
}
