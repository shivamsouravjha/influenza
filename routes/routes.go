package routes

import (
	"github.com/gin-gonic/gin"
	get "github.com/shivamsouravjha/influenza/controllers/GET"
)

func NewRouter() *gin.Engine {

	router := gin.New()
	v1 := router.Group("/api")

	emailRoutes := v1.Group("/email")
	{
		emailRoutes.GET("/getVerificationCode", get.VerificationCode)
	}
	return router
}
