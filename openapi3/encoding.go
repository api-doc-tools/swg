package openapi3

// Encoding is specified by OpenAPI/Swagger 3.0 standard.
type Encoding struct {
	ExtensionProps

	ContentType   string                `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	Headers       map[string]*HeaderRef `json:"headers,omitempty" yaml:"headers,omitempty"`
	Style         string                `json:"style,omitempty" yaml:"style,omitempty"`
	Explode       bool                  `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved bool                  `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
}
