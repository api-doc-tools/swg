package openapi3

type CallbackRef struct {
	Ref   string
	Value *Callback
}

type ExampleRef struct {
	Ref   string
	Value *Example
}

type HeaderRef struct {
	Ref   string
	Value *Header
}

type LinkRef struct {
	Ref   string
	Value *Link
}

type ParameterRef struct {
	Ref   string
	Value *Parameter
}

type ResponseRef struct {
	Ref   string
	Value *Response
}

type RequestBodyRef struct {
	Ref   string
	Value *RequestBody
}

type SchemaRef struct {
	Ref   string
	Value *Schema
}

type SecuritySchemeRef struct {
	Ref   string
	Value *SecurityScheme
}
