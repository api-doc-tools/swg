package openapi3

// Link is specified by OpenAPI/Swagger standard version 3.0.
type Link struct {
	ExtensionProps
	Description string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Href        string                 `json:"href,omitempty" yaml:"href,omitempty"`
	OperationID string                 `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Parameters  map[string]interface{} `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Headers     map[string]*Schema     `json:"headers,omitempty" yaml:"headers,omitempty"`
}
