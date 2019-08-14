package controllers

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

const (
	elasticIndexName = "facebook-group"
	elasticTypeName  = "document"
)

func SearchEndpoint(c *gin.Context) {
	// Declarasi struct
	type SearchResponse = models.SearchResponse
	type DocumentResponse = models.DocumentResponse

	// Create Elastic client and wait for Elasticsearch to be ready
	elasticClient := config.DBInit()

	// Parse request
	query := c.Query("query")
	if query == "" {
		errorResponse(c, http.StatusBadRequest, "Query not specified")
		return
	}
	skip := 0
	take := 10
	if i, err := strconv.Atoi(c.Query("skip")); err == nil {
		skip = i
	}
	if i, err := strconv.Atoi(c.Query("take")); err == nil {
		take = i
	}
	// Perform search
	esQuery := elastic.NewMatchQuery("province_single", query).
		Fuzziness("2").
		MinimumShouldMatch("2")
	result, err := elasticClient.Search().
		Index(elasticIndexName).
		Query(esQuery).
		From(skip).Size(take).
		Do(c.Request.Context())
	if err != nil {
		log.Println(err)
		errorResponse(c, http.StatusInternalServerError, "Something went wrong")
		return
	}
	res := SearchResponse{
		Time: fmt.Sprintf("%d", result.TookInMillis),
		Hits: fmt.Sprintf("%d", result.Hits.TotalHits),
	}

	// Transform search results before returning them
	docs := make([]DocumentResponse, 0)
	for _, hit := range result.Hits.Hits {
		var doc DocumentResponse
		if err := json.Unmarshal(*hit.Source, &doc); err != nil {
			panic(err)
		}
		docs = append(docs, doc)
	}
	res.Documents = docs
	c.JSON(http.StatusOK, res)
}

func CreateDocumentsEndpoint(c *gin.Context) {
	var (
		docs []models.DocumentRequest
	)
	type Document = models.Document

	// Create Elastic client and wait for Elasticsearch to be ready
	elasticClient := config.DBInit()

	// Parse request
	if err := c.BindJSON(&docs); err != nil {
		errorResponse(c, http.StatusBadRequest, "Malformed request body")
		return
	}
	// Insert documents in bulk
	bulk := elasticClient.
		Bulk().
		Index(elasticIndexName).
		Type(elasticTypeName)
	for _, d := range docs {
		doc := Document{
			ID:        shortid.MustGenerate(),
			Title:     d.Title,
			CreatedAt: time.Now().UTC(),
			Content:   d.Content,
		}
		bulk.Add(elastic.NewBulkIndexRequest().Id(doc.ID).Doc(doc))
	}
	if _, err := bulk.Do(c.Request.Context()); err != nil {
		log.Println(err)
		errorResponse(c, http.StatusInternalServerError, "Failed to create documents")
		return
	}
	c.Status(http.StatusOK)
}

func errorResponse(c *gin.Context, code int, err string) {
	c.JSON(code, gin.H{
		"error": err,
	})
}
