package config

import (
	"log"
	"time"

	"github.com/olivere/elastic"
)

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
