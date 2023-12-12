package routes

import (
	"hrms/controllers"
	"hrms/models"

	"github.com/gin-gonic/gin"
)

func Init() {

}

func Routes() {

	r := gin.Default()
	// Use CORS middleware
	//config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"*"}
	//r.Use(cors.New(config))
	r.Use(CORSMiddleware())

	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	models.Connect()

	r.Static("/public", "./public")

	r.POST("/login", controllers.Login)

	r.GET("/users/:limit/:offset", controllers.IndexUsers)
	r.POST("/users/store", controllers.StoreUsers)
	r.GET("/users/show/:id", controllers.ShowUser)

	r.GET("/documents/:limit/:offset", controllers.IndexDocuments)
	r.GET("/searchByTicketNo/:searchTerm", controllers.SearchByTicketNo)
	r.GET("/user/documents/:user_id", controllers.FetchByUserId)
	r.POST("/documents/store/:userid", controllers.StoreDocuments)
	r.POST("/add-documents/:name/:user_id", controllers.AddDocument)
	r.PATCH("/documents/update/:id/:name", controllers.UpdateDocument)
	r.Run(":8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
