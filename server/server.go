package server

import "github.com/shivamsouravjha/influenza/routes"

func Init() {
	r := routes.NewRouter()
	r.Run(":8080")
}
