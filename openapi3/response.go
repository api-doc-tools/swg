package openapi3

// Responses is specified by OpenAPI/Swagger 3.0 standard.
type Responses map[string]*ResponseRef

// Response is specified by OpenAPI/Swagger 3.0 standard.
type Response struct {
	ExtensionProps
	Description string              `json:"description,omitempty" yaml:"description,omitempty"`
	Headers     map[string]*Schema  `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     Content             `json:"content,omitempty" yaml:"content,omitempty"`
	Links       map[string]*LinkRef `json:"links,omitempty" yaml:"links,omitempty"`
}
