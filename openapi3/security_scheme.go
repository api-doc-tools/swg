package openapi3

type SecurityScheme struct {
	ExtensionProps

	Type         string      `json:"type,omitempty" yaml:"type,omitempty"`
	Description  string      `json:"description,omitempty" yaml:"description,omitempty"`
	Name         string      `json:"name,omitempty" yaml:"name,omitempty"`
	In           string      `json:"in,omitempty" yaml:"in,omitempty"`
	Scheme       string      `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	BearerFormat string      `json:"bearerFormat,omitempty" yaml:"bearerFormat,omitempty"`
	Flows        *OAuthFlows `json:"flows,omitempty" yaml:"flows,omitempty"`
}

type OAuthFlows struct {
	ExtensionProps
	Implicit          *OAuthFlow `json:"implicit,omitempty" yaml:"implicit,omitempty"`
	Password          *OAuthFlow `json:"password,omitempty" yaml:"password,omitempty"`
	ClientCredentials *OAuthFlow `json:"clientCredentials,omitempty" yaml:"clientCredentials,omitempty"`
	AuthorizationCode *OAuthFlow `json:"authorizationCode,omitempty" yaml:"authorizationCode,omitempty"`
}

type oAuthFlowType int

const (
	oAuthFlowTypeImplicit oAuthFlowType = iota
	oAuthFlowTypePassword
	oAuthFlowTypeClientCredentials
	oAuthFlowAuthorizationCode
)

type OAuthFlow struct {
	ExtensionProps
	AuthorizationURL string            `json:"authorizationUrl,omitempty" yaml:"description,omitempty" `
	TokenURL         string            `json:"tokenUrl,omitempty" yaml:"tokenUrl,omitempty"`
	RefreshURL       string            `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`
	Scopes           map[string]string `json:"scopes" yaml:"scopes"`
}
