package openapi3

// ExtensionProps provides support for OpenAPI extensions.
// It reads/writes all properties that begin with "x-".
type ExtensionProps struct {
	Extensions map[string]interface{} `json:"-"`
}
