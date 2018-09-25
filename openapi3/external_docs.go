package openapi3

// ExternalDocs is specified by OpenAPI/Swagger standard version 3.0.
type ExternalDocs struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	URL         string `json:"url,omitempty" yaml:"url,omitempty"`
}
