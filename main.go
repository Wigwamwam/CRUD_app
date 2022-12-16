package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wigwamwam/CRUD_app/controllers"
	"github.com/wigwamwam/CRUD_app/initializers"
)

func init() {
	initializers.LoadEnvVairables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.POST("/banks", controllers.BanksCreate)
	r.PUT("/banks/:id", controllers.BanksUpdate)
	r.GET("/banks", controllers.BanksIndex)
	r.GET("/banks/:id", controllers.BanksShow)
	r.DELETE("/banks/:id", controllers.BanksDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
}
