package handlers

import (
	"dbpedia-server/interfaces"
	"dbpedia-server/types"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Handler contains all HTTP handlers
type Handler struct {
	dbpediaClient interfaces.DBpediaClient
}

// NewHandler creates a new handler instance
func NewHandler(client interfaces.DBpediaClient) *Handler {
	return &Handler{
		dbpediaClient: client,
	}
}

// HealthCheck returns the server status
func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, types.HealthResponse{
		Status: "healthy",
		Time:   time.Now().Format(time.RFC3339),
	})
}

// ExecuteSPARQLQuery executes a SPARQL query against DBpedia
func (h *Handler) ExecuteSPARQLQuery(c *gin.Context) {
	var req types.SPARQLRequest

	// Bind JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error:   "invalid_request",
			Message: "Query parameter is required",
		})
		return
	}

	// Call DBpedia SPARQL endpoint
	result, err := h.dbpediaClient.Query(req.Query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:   "query_failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetExampleQueries returns some example SPARQL queries
func (h *Handler) GetExampleQueries(c *gin.Context) {
	examples := []types.ExampleQuery{
		{
			Name:        "Get information about Albert Einstein",
			Description: "Retrieve basic information about Albert Einstein from DBpedia",
			Query: `SELECT ?property ?value WHERE {
  <http://dbpedia.org/resource/Albert_Einstein> ?property ?value .
} LIMIT 10`,
		},
		{
			Name:        "List 10 scientists",
			Description: "Get a list of 10 scientists from DBpedia",
			Query: `SELECT ?scientist ?name WHERE {
  ?scientist a <http://dbpedia.org/ontology/Scientist> .
  ?scientist <http://xmlns.com/foaf/0.1/name> ?name .
} LIMIT 10`,
		},
		{
			Name:        "Count entities by type",
			Description: "Count the number of different types of entities",
			Query: `SELECT ?type (COUNT(?s) as ?count) WHERE {
  ?s a ?type .
} GROUP BY ?type ORDER BY DESC(?count) LIMIT 10`,
		},
		{
			Name:        "Get countries and capitals",
			Description: "Retrieve countries with their capital cities",
			Query: `SELECT ?country ?countryName ?capital ?capitalName WHERE {
  ?country a <http://dbpedia.org/ontology/Country> .
  ?country <http://xmlns.com/foaf/0.1/name> ?countryName .
  ?country <http://dbpedia.org/ontology/capital> ?capital .
  ?capital <http://xmlns.com/foaf/0.1/name> ?capitalName .
  FILTER (lang(?countryName) = 'en' && lang(?capitalName) = 'en')
} LIMIT 10`,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"examples": examples,
	})
}
