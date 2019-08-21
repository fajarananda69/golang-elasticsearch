# golang-elasticsearch

## Install and SETUP GO
Download and configure your workspace with latest version of Go and correct environment path.
- [Last Version](https://golang.org/dl/)
- [Windows](http://www.wadewegner.com/2014/12/easy-go-programming-setup-for-windows/)
- [Linux](http://www.tecmint.com/install-go-in-linux/)

## Install Driver
Install dep [klik here](https://golang.github.io/dep/docs/installation.html) and Create new dep
```
dep init 
```
Install driver to build rest and elastic
```
dep ensure -add github.com/gin-gonic/gin
```
```
dep ensure -add github.com/olivere/elastic
```

Import driver
```
import (
	"encoding/json"
	"fmt"
	"go-elasticsearch/search-api/config"
	"go-elasticsearch/search-api/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
	"github.com/teris-io/shortid"
)
```

## Connection
```
func DBInit() *elastic.Client {
	var (
		err           error
		elasticClient *elastic.Client
	)
	// Create Elastic client and wait for Elasticsearch to be ready
	for {
		elasticClient, err = elastic.NewClient(
			elastic.SetURL("http://localhost:9200"),
			elastic.SetSniff(false),
		)
		if err != nil {
			log.Println(err)
			// Retry every 3 seconds
			time.Sleep(3 * time.Second)
		} else {
			break
		}
	}
	return elasticClient
}
```

## gin controllers
```
var err error
r := gin.Default()
	r.POST("/documents", controllers.CreateDocumentsEndpoint)
	r.GET("/search", controllers.SearchEndpoint)
	if err = r.Run(":9999"); err != nil {
		log.Fatal(err)
	}
```
build rest with method POST and GET, if you run this file it will be run on port 9999 `r.run(:9999)`
## Run
```
go run main.go
```
