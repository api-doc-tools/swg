package openapi3

// Components is specified by OpenAPI/Swagger standard version 3.0.
type Components struct {
	ExtensionProps
	Schemas         map[string]*SchemaRef         `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	Parameters      map[string]*ParameterRef      `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Headers         map[string]*HeaderRef         `json:"headers,omitempty" yaml:"headers,omitempty"`
	RequestBodies   map[string]*RequestBodyRef    `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
	Responses       map[string]*ResponseRef       `json:"responses,omitempty" yaml:"responses,omitempty"`
	SecuritySchemes map[string]*SecuritySchemeRef `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`
	Examples        map[string]*ExampleRef        `json:"examples,omitempty" yaml:"examples,omitempty"`
	Tags            Tags                          `json:"tags,omitempty" yaml:"tags,omitempty"`
	Links           map[string]*LinkRef           `json:"links,omitempty" yaml:"links,omitempty"`
	Callbacks       map[string]*CallbackRef       `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
}
