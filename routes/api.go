package routes

import (
	"hrms/controllers"
	"hrms/models"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()

	models.Connect()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"result": "This is the result"})
	})
	r.GET("/users/:offset", controllers.Index)
	r.Run(":8080")
}
