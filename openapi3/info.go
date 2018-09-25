package openapi3

// Info is specified by OpenAPI/Swagger standard version 3.0.
type Info struct {
	ExtensionProps
	Title          string   `json:"title,omitempty" yaml:"title,omitempty"`
	Description    string   `json:"description,omitempty" yaml:"description,omitempty"`
	TermsOfService string   `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *License `json:"license,omitempty" yaml:"license,omitempty"`
	Version        string   `json:"version,omitempty" yaml:"version,omitempty"`
}

// Contact is specified by OpenAPI/Swagger standard version 3.0.
type Contact struct {
	ExtensionProps
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	URL   string `json:"url,omitempty" yaml:"url,omitempty"`
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}

// License is specified by OpenAPI/Swagger standard version 3.0.
type License struct {
	ExtensionProps
	Name string `json:"name,omitempty" yaml:"email,omitempty"`
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}
