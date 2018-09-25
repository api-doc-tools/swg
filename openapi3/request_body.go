package openapi3

// RequestBody is specified by OpenAPI/Swagger 3.0 standard.
type RequestBody struct {
	ExtensionProps
	Description string  `json:"description,omitempty" yaml:"description,omitempty"`
	Required    bool    `json:"required,omitempty" yaml:"required,omitempty"`
	Content     Content `json:"content,omitempty" yaml:"content,omitempty"`
}
