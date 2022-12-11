package main

import (
	"demo-api/controllers"
	"demo-api/initializers"

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
	r.GET("/users", controllers.GetUsers)
	r.POST("/login", controllers.Login)
	r.Run() // listen and serve on 0.0.0.0:8080
}
