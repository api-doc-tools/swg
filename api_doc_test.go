package swg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type BookImageUrls struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"larger"`
}

type Book struct {
	ID        string        `json:"id"`
	Title     string        `json:"title"`
	Summary   string        `json:"summary"`
	Authors   []string      `json:"authors"`
	Images    BookImageUrls `json:"images"`
	Pages     int           `json:"pages"`
	Price     float32       `json:"price"`
	HasReview bool          `json:"has_review"`
}

type errorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"detail"`
}
type bookImageUrls struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"larger"`
}
type book struct {
	ID        string        `json:"id"`
	Title     string        `json:"title"`
	Summary   string        `json:"summary"`
	Authors   []string      `json:"authors"`
	Images    bookImageUrls `json:"images"`
	Pages     int           `json:"pages"`
	Price     float32       `json:"price"`
	HasReview bool          `json:"has_review"`
}

type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"detail"`
}

var docGETBook = &APIDocCommon{
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

var invalidDocGETBook = &APIDocCommon{
	Summary:  "Get book info by id",
	Produces: []string{Application_Json},
	Parameters: map[string]Parameter{
		"id":      Parameter{InPath: &Value{Type: "string"}},
		"version": Parameter{InHeader: &Value{Type: "string", Desc: "the version of api"}},
		"range":   Parameter{InFormData: &Value{Type: "string"}}, // err:  FormData is not supported in method GET
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

var invalidDocGETBook2 = &APIDocCommon{
	Summary:  "Get book info by id",
	Produces: []string{Application_Json},
	Parameters: map[string]Parameter{
		"id":      Parameter{InPath: &Value{Type: "string"}},
		"version": Parameter{InHeader: &Value{Type: "string", Desc: "the version of api"}},
	},
	Request: &Request{ // err: in method GET, Request should be nil
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

var docPostBook = &APIDocCommon{
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

var invalidDocPostBook = &APIDocCommon{
	Summary:  "new a book",
	Produces: []string{Application_Json},
	Consumes: []string{Application_Json},
	Parameters: map[string]Parameter{
		"version": Parameter{InHeader: &Value{Type: "string", Desc: "the version of api"}},
		// err:  Parameter.InFormData and APIDocCommon.Request can't coexist
		// There are parameters in formData, doc.Request should be nil
		"range": Parameter{InFormData: &Value{Type: "string", Desc: "range"}},
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

var invalidDocPostBook2 = &APIDocCommon{
	Summary:  "new a book",
	Produces: []string{Application_Json},
	Consumes: []string{Application_Json},
	Parameters: map[string]Parameter{
		"version": Parameter{InHeader: &Value{Type: "string", Desc: "the version of api"}},
	},
	Request: &Request{
		Description: "the book info",
		Model:       nil, //err : doc.Request should not be nil
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

func TestAPIDocCommon_ToSwaggerOperation(t *testing.T) {
	assert := assert.New(t)
	// get
	docGETBook.SetMethod(GET)
	_, err := docGETBook.ToSwaggerOperation()
	assert.Nil(err)

	invalidDocGETBook.SetMethod(GET)
	_, err = invalidDocGETBook.ToSwaggerOperation()
	assert.NotNil(err)

	invalidDocGETBook2.SetMethod(GET)

	_, err = invalidDocGETBook2.ToSwaggerOperation()
	assert.NotNil(err)

	// post
	docPostBook.SetMethod(POST)
	_, err = docPostBook.ToSwaggerOperation()
	assert.Nil(err)

	invalidDocPostBook.SetMethod(POST)
	_, err = invalidDocPostBook.ToSwaggerOperation()
	assert.NotNil(err)

	_, err = invalidDocPostBook2.ToSwaggerOperation()
	assert.NotNil(err)
}

func TestAPIDocCommon_ToSwaggerDefinitions(t *testing.T) {
	assert := assert.New(t)
	_, err := docGETBook.ToSwaggerDefinitions()
	assert.Nil(err)

	_, err = docPostBook.ToSwaggerDefinitions()
	assert.Nil(err)
}
