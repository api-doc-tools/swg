package openapi3

// Float64Ptr is a helper for defining OpenAPI schemas.
func Float64Ptr(value float64) *float64 {
	return &value
}

// BoolPtr is a helper for defining OpenAPI schemas.
func BoolPtr(value bool) *bool {
	return &value
}

// Int64Ptr is a helper for defining OpenAPI schemas.
func Int64Ptr(value int64) *int64 {
	return &value
}

// Uint64Ptr is a helper for defining OpenAPI schemas.
func Uint64Ptr(value uint64) *uint64 {
	return &value
}

// Schema is specified by OpenAPI/Swagger 3.0 standard.
type Schema struct {
	ExtensionProps

	OneOf        []*SchemaRef  `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`
	AnyOf        []*SchemaRef  `json:"anyOf,omitempty" yaml:"anyOf,omitempty"`
	AllOf        []*SchemaRef  `json:"allOf,omitempty" yaml:"allOf,omitempty"`
	Not          *SchemaRef    `json:"not,omitempty" yaml:"not,omitempty"`
	Type         string        `json:"type,omitempty" yaml:"type,omitempty"`
	Format       string        `json:"format,omitempty" yaml:"format,omitempty"`
	Description  string        `json:"description,omitempty" yaml:"description,omitempty"`
	Enum         []interface{} `json:"enum,omitempty" yaml:"enum,omitempty"`
	Default      interface{}   `json:"default,omitempty" yaml:"default,omitempty"`
	Example      interface{}   `json:"example,omitempty" yaml:"example,omitempty"`
	ExternalDocs interface{}   `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`

	// Object-related, here for struct compactness
	AdditionalPropertiesAllowed *bool `json:"-" multijson:"additionalProperties,omitempty" yaml:"-"`
	// Array-related, here for struct compactness
	UniqueItems bool `json:"uniqueItems,omitempty" yaml:"uniqueItems,omitempty"`
	// Number-related, here for struct compactness
	ExclusiveMin bool `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`
	ExclusiveMax bool `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`
	// Properties
	Nullable  bool        `json:"nullable,omitempty" yaml:"nullable,omitempty" `
	ReadOnly  bool        `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	WriteOnly bool        `json:"writeOnly,omitempty" yaml:"writeOnly,omitempty"`
	XML       interface{} `json:"xml,omitempty" yaml:"xml,omitempty"`

	// Number
	Min        *float64 `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Max        *float64 `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	MultipleOf *float64 `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`

	// String
	MinLength uint64  `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	MaxLength *uint64 `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	Pattern   string  `json:"pattern,omitempty" yaml:"pattern,omitempty"`

	// Array
	MinItems uint64     `json:"minItems,omitempty" yaml:"minItems,omitempty"`
	MaxItems *uint64    `json:"maxItems,omitempty" yaml:"maxItems,omitempty"`
	Items    *SchemaRef `json:"items,omitempty" yaml:"items,omitempty"`

	// Object
	Required             []string              `json:"required,omitempty" yaml:"required,omitempty"`
	Properties           map[string]*SchemaRef `json:"properties,omitempty" yaml:"properties,omitempty"`
	MinProps             uint64                `json:"minProperties,omitempty" yaml:"minProperties,omitempty"`
	MaxProps             *uint64               `json:"maxProperties,omitempty" yaml:"maxProperties,omitempty"`
	AdditionalProperties *SchemaRef            `json:"-" multijson:"additionalProperties,omitempty" yaml:"-"`
	Discriminator        string                `json:"discriminator,omitempty" yaml:"discriminator,omitempty"`

	PatternProperties string `json:"patternProperties,omitempty" yaml:"patternProperties,omitempty"`
}
