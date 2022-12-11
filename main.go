package main

import (
	"demo-api/controllers"
	"demo-api/initializers"
	"demo-api/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
	initializers.MigrateDB()
}

func main() {
	r := gin.Default()
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", middleware.Auth, controllers.GetUsers)
	r.GET("/users/:id", middleware.Auth, controllers.GetUser)
	r.POST("/login", controllers.Login)
	r.Run() // listen and serve on 0.0.0.0:8080
}
