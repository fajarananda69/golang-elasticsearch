package main

import (
	"go-elasticsearch/search-api/controllers"
	"go-elasticsearch/search-api/docs"
	"log"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	var err error

	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Start HTTP server
	r := gin.Default()
	// url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.POST("/documents", controllers.CreateDocumentsEndpoint)
	r.GET("/search", controllers.SearchEndpoint)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if err = r.Run(":9999"); err != nil {
		log.Fatal(err)
	}
}
