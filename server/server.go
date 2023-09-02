package server

import "github.com/gin-gonic/gin"

func Init() {
	r := gin.New()
	r.Run(":8080")
}
