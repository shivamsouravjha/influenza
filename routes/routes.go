package routes

import (
	"github.com/gin-gonic/gin"
	get "github.com/shivamsouravjha/influenza/controllers/GET"
	post "github.com/shivamsouravjha/influenza/controllers/POST"
	"github.com/shivamsouravjha/influenza/helpers/redis"
)

func NewRouter() *gin.Engine {
	redis.RedisInit()
	router := gin.New()
	v1 := router.Group("/api")

	emailRoutes := v1.Group("/email")
	{
		emailRoutes.GET("/getVerificationCode", get.VerificationCode)
		emailRoutes.POST("/verifyCode", post.VerifyCode)
	}
	return router
}
