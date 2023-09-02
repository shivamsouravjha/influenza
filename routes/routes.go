package routes

import (
	"github.com/gin-gonic/gin"
	get "github.com/shivamsouravjha/influenza/controllers/GET"
	"github.com/shivamsouravjha/influenza/helpers"
)

func NewRouter() *gin.Engine {
	helpers.RedisInit()
	router := gin.New()
	v1 := router.Group("/api")

	emailRoutes := v1.Group("/email")
	{
		emailRoutes.GET("/getVerificationCode", get.VerificationCode)
	}
	return router
}
