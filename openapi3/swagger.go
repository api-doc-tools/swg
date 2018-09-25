package openapi3

type Swagger struct {
	ExtensionProps
	OpenAPI      string               `json:"openapi" yaml:"openapi"` // Required
	Info         Info                 `json:"info" yaml:"info"`       // Required
	Servers      Servers              `json:"servers,omitempty" yaml:"servers,omitempty"`
	Paths        Paths                `json:"paths,omitempty" yaml:"paths,omitempty"`
	Components   Components           `json:"components,omitempty" yaml:"components,omitempty"`
	Security     SecurityRequirements `json:"security,omitempty" yaml:"security,omitempty"`
	ExternalDocs *ExternalDocs        `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}
