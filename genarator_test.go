package swg

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"
)

type Books struct {
	Total int64   `json:"total" desc:"total of zoos"`
	Start int64   `json:"start"`
	Count int64   `json:"count"`
	Books []*Book `json:"books" desc:"books"`
}

var DocGETBook = &APIDocCommon{
	Summary:  "Get book info by id",
	Produces: []string{Application_Json},
	Parameters: map[string]Parameter{
		"id":      Parameter{InPath: &Value{Type: "string"}},
		"version": Parameter{InHeader: &Value{Type: "string", Desc: "the version of api"}},
	},
	Responses: map[int]Response{
		200: Response{
			Description: "successful operation",
			Model:       &Book{},
		},
		400: Response{
			Description: "failed operation",
			Model:       &ErrorMessage{},
		},
	},
}

var DocPostBook = &APIDocCommon{
	Summary:  "new a book",
	Produces: []string{Application_Json},
	Consumes: []string{Application_Json},
	Parameters: map[string]Parameter{
		"version": Parameter{InHeader: &Value{Type: "string", Desc: "the version of api"}},
	},
	Request: &Request{
		Description: "the book info",
		Model:       &Book{},
	},
	Responses: map[int]Response{
		200: Response{
			Description: "successful operation",
			Model:       &Book{},
		},
		400: Response{
			Description: "failed operation",
			Model:       &ErrorMessage{},
		},
	},
}

var DocDELETEBook = &APIDocCommon{
	Summary:  "delete book info by id",
	Produces: []string{Application_Json},
	Parameters: map[string]Parameter{
		"id":      Parameter{InPath: &Value{Type: "string"}},
		"version": Parameter{InHeader: &Value{Type: "string", Desc: "the version of api"}},
	},
	Responses: map[int]Response{
		200: Response{
			Description: "successful operation",
			Model:       &Book{},
		},
		400: Response{
			Description: "failed operation",
			Model:       &ErrorMessage{},
		},
	},
}

var DocGETBooks = &APIDocCommon{
	Summary:  "Get book info by id",
	Produces: []string{Application_Json},
	Parameters: map[string]Parameter{
		"version":   Parameter{InHeader: &Value{Type: "string", Desc: "the version of api"}},
		"limit":     Parameter{InQuery: &Value{Type: "int64", Min: "0", Max: "1000", Required: true, Desc: "the limit of searching"}},
		"offset":    Parameter{InQuery: &Value{Type: "int64", Required: true, Desc: "the offset of searching"}},
		"sort":      Parameter{InQuery: &Value{Type: "string", Enum: "id -id price -price", Desc: "sort of searching"}},
		"min_price": Parameter{InQuery: &Value{Type: "float32", Min: "0", Desc: "minimum price"}},
		"max_price": Parameter{InQuery: &Value{Type: "float32", Min: "0", Desc: "minimum price"}},
	},
	Responses: map[int]Response{
		200: Response{
			Description: "successful operation",
			Model:       &Books{},
		},
		400: Response{
			Description: "failed operation",
			Model:       &ErrorMessage{},
		},
	},
}

var conf = &Config{
	Schemes:     []Scheme{SchemeHTTP, SchemeHTTPS},
	BasePath:    "/dev",
	Version:     "v1",
	Title:       "Demo APIS",
	Description: "Description of Demo APIS",
}

func TestGenarator(t *testing.T) {
	assert := assert.New(t)
	router := NewDocGenarator(conf)
	router.GET("/book/:id/", DocGETBook)
	router.GET("/books", DocGETBooks)
	router.POST("/books", DocPostBook)

	yamlBytes, err := yaml.Marshal(router.Swagger)
	assert.Nil(err)
	err = ioutil.WriteFile("swagger.yaml", yamlBytes, 0644)
	assert.Nil(err)
}

func TestInvalidDoc(t *testing.T) {
	missParamterInDoc(t)
}

func missParamterInDoc(t *testing.T) {
	assert := assert.New(t)
	defer func() {
		if r := recover(); r != nil {
			assert.NotNil(r)
		}
	}()
	router := NewDocGenarator(conf)
	router.GET("/books/:id", &APIDocCommon{
		Summary:  "Get book info by id",
		Produces: []string{Application_Json},
		Responses: map[int]Response{
			200: Response{
				Description: "successful operation",
				Model:       &Book{},
			},
			400: Response{
				Description: "failed operation",
				Model:       &ErrorMessage{},
			},
		},
	})
}

func TestEngine_GetSwaggerJSONDocument(t *testing.T) {
	assert := assert.New(t)
	router := NewDocGenarator(conf)
	_, err := router.GetSwaggerJSONDocument()
	assert.Nil(err)
}

func TestEngine_GetSwaggerYAMLDocument(t *testing.T) {
	assert := assert.New(t)
	router := NewDocGenarator(conf)
	_, err := router.GetSwaggerYAMLDocument()
	assert.Nil(err)
}
