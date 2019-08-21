# golang-elasticsearch

## Install Driver
Create new dep
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
## Run
```
go run main.go
```
