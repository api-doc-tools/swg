package openapi3

// Servers is specified by OpenAPI/Swagger standard version 3.0.
type Servers []*Server

// Server is specified by OpenAPI/Swagger standard version 3.0.
type Server struct {
	URL         string                     `json:"url,omitempty" yaml:"url,omitempty"`
	Description string                     `json:"description,omitempty" yaml:"description,omitempty"`
	Variables   map[string]*ServerVariable `json:"variables,omitempty" yaml:"variables,omitempty"`
}

// ServerVariable is specified by OpenAPI/Swagger standard version 3.0.
type ServerVariable struct {
	Enum        []interface{} `json:"enum,omitempty" yaml:"enum,omitempty"`
	Default     interface{}   `json:"default,omitempty" yaml:"default,omitempty"`
	Description string        `json:"description,omitempty" yaml:"description,omitempty"`
}
