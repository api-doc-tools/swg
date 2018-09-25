package openapi3

type Header struct {
	ExtensionProps

	// Optional description. Should use CommonMark syntax.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Optional schema
	Schema *SchemaRef `json:"schema,omitempty" yaml:"schema,omitempty"`
}
