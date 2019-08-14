package main

import (
	"go-elasticsearch/search-api/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	var err error

	// Start HTTP server
	r := gin.Default()
	r.POST("/documents", controllers.CreateDocumentsEndpoint)
	r.GET("/search", controllers.SearchEndpoint)
	if err = r.Run(":9999"); err != nil {
		log.Fatal(err)
	}
}
