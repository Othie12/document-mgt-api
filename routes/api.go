package routes

import (
	"hrms/controllers"
	"hrms/models"
	"os"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()

	models.Connect()

	p := os.Getenv("UPLOAD_PATH")

	r.Static("/public", p)
	/*
		r.GET("/files/:filename", func(c *gin.Context) {
			filename := c.Param("filename")
			filePath := filepath.Join(p, filename)
			c.File(filePath)
		})
	*/
	r.POST("login", controllers.Login)

	r.GET("/users/:limit/:offset", controllers.IndexUsers)
	r.POST("/users/store", controllers.StoreUsers)

	r.GET("/documents/:limit/:offset", controllers.IndexDocuments)
	r.GET("/searchByTicketNo/:searchTerm", controllers.SearchByTicketNo)
	r.POST("/documents/store/:userid", controllers.StoreDocuments)
	r.POST("/add-documents/:name/:userId", controllers.AddDocument)
	r.Run(":8080")
}
