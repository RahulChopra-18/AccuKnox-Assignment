package main

import (
	controller "go-crud/controllers"
	"go-crud/initializers"
	"go-crud/middleware"

	"github.com/gin-gonic/gin"
)

func init() { //it will run before main
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {

	r := gin.Default()

	r.POST("/signup", controller.Signup)
	r.POST("/login", controller.Login)

	r.POST("/notes", middleware.RequireAuth, controller.NotesCreate)
	r.PUT("/notes/:id", middleware.RequireAuth, controller.NotesUpdate)
	r.GET("/notes", middleware.RequireAuth, controller.NotesIndex)
	r.GET("/notes/:id", middleware.RequireAuth, controller.NotesShow)
	r.DELETE("/notes/:id", middleware.RequireAuth, controller.NotesDelete)
	r.Run()
}
