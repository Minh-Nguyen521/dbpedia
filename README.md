# DBpedia

## 📚 What is DBpedia?

**DBpedia** is a crowd-sourced community effort to extract structured information from Wikipedia and make it available as a knowledge graph on the web.

### Key Facts About DBpedia

- **Scope**: Contains structured data from over 125 Wikipedia language editions
- **Scale**: Describes more than 6 million entities (people, places, organizations, etc.)
- **Interconnected**: Links to other datasets like Wikidata, GeoNames, and more
- **Free & Open**: Available under open licenses (CC BY-SA and GFDL)
- **Real-world Use**: Powers semantic search, chatbots, recommendation systems, and AI applications

## SPARQL

### 1. Execute SPARQL Query
**POST** `/api/v1/sparql` or `/sparql`

Execute a SPARQL query against DBpedia.

**Request:**
```bash
curl -X POST http://localhost:8080/api/v1/sparql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "SELECT ?name WHERE { <http://dbpedia.org/resource/Albert_Einstein> <http://xmlns.com/foaf/0.1/name> ?name . } LIMIT 1"
  }'
```

**Response:**
```json
{
  "head": {
    "vars": ["name"]
  },
  "results": {
    "bindings": [
      {
        "name": {
          "type": "literal",
          "value": "Albert Einstein",
          "xml:lang": "en"
        }
      }
    ]
  }
}
```

---

### 2. Validate SPARQL Query
**POST** `/api/v1/validate`

Validate a SPARQL query without executing it.

**Request:**
```bash
curl -X POST http://localhost:8080/api/v1/validate \
  -H "Content-Type: application/json" \
  -d '{
    "query": "SELECT ?s ?p ?o WHERE { ?s ?p ?o . } LIMIT 10"
  }'
```

**Response:**
```json
{
  "valid": true,
  "message": "Query is valid"
}
```

## Basic SPARQL Query Structure

```sparql
SELECT ?variable1 ?variable2
WHERE {
  ?subject ?predicate ?object .
  # More triple patterns...
}
LIMIT 10
```

### Key Concepts

1. **Triple Pattern**: `?subject ?predicate ?object`
   - Subject: The resource you're querying
   - Predicate: The property/relationship
   - Object: The value or related resource

2. **Variables**: Start with `?` (e.g., `?name`, `?country`)

3. **URIs**: Resources in DBpedia (e.g., `<http://dbpedia.org/resource/Albert_Einstein>`)

### Example

#### 1. Get Information About a Specific Resource

```sparql
SELECT ?property ?value WHERE {
  <http://dbpedia.org/resource/Albert_Einstein> ?property ?value .
} LIMIT 10
```

Use the `/api/v1/sparql` endpoint with this query in your request body.

#### 2. Find Resources by Type

```sparql
SELECT ?scientist ?name WHERE {
  ?scientist a <http://dbpedia.org/ontology/Scientist> .
  ?scientist <http://xmlns.com/foaf/0.1/name> ?name .
} LIMIT 10
```

#### 3. Filter Results

```sparql
SELECT ?country ?name WHERE {
  ?country a <http://dbpedia.org/ontology/Country> .
  ?country <http://xmlns.com/foaf/0.1/name> ?name .
  FILTER (lang(?name) = 'en')
} LIMIT 10
```

#### 4. Aggregate Data

```sparql
SELECT ?type (COUNT(?s) as ?count) WHERE {
  ?s a ?type .
} GROUP BY ?type ORDER BY DESC(?count) LIMIT 10
```

#### 5. Simple Integration (cURL)

```bash
# Get information about Paris
curl -X POST http://localhost:8080/api/v1/sparql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "SELECT ?abstract WHERE { <http://dbpedia.org/resource/Paris> <http://dbpedia.org/ontology/abstract> ?abstract . FILTER (lang(?abstract) = \"en\") } LIMIT 1"
  }'
```

## DBPEDIA REPO

https://github.com/aniiyengar/dbpedia

https://github.com/garfix/nli-go

# THIS CURRENT PROJECT'S STRUCTURE

1. **main.go**: Application entry point
   - Loads configuration
   - Initializes DBpedia client
   - Sets up HTTP server
   - **Don't modify** unless changing app initialization

2. **client/dbpedia_client.go**: DBpedia SPARQL client
   - Call `NewDBpediaClient()` to create a client instance
   - Call `Query()` method to execute SPARQL queries
   - **Modify** if you need to change how queries are sent to DBpedia

3. **handlers/handlers.go**: HTTP request handlers
   - Contains `ExecuteSPARQLQuery`, `ValidateSPARQLQuery`, `GetExampleQueries`
   - **Modify** to add new endpoints or change request/response handling

4. **server/server.go**: HTTP server configuration
   - Call `SetupRoutes()` to define API endpoints
   - **Modify** to add new routes or middleware

5. **types/types.go**: Data structures
   - Defines request/response types
   - **Modify** when adding new API endpoints or changing data structures