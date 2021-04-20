package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"GO14/assignment02/controllers"
	"GO14/assignment02/config"
)

func Init() {
	db := config.Db
	
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "berhasil",
		})
	})

	r.GET("/orders", controllers.GetOrders)
	r.POST("/orders", controllers.CreateOrders)
	r.PUT("/orders/:id", controllers.EditOrders)
	r.DELETE("/orders/:id", controllers.DeleteOrders)

	fmt.Println("starting web server at http://localhost:8080/")
	r.Run()
}