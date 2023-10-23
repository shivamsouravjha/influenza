package server

import (
	"github.com/shivamsouravjha/influenza/config"
	"github.com/shivamsouravjha/influenza/routes"
)

func Init() {
	r := routes.NewRouter()
	r.Run(":" + config.Get().ServerPort) //running the server at port
}
