package swg

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/api-doc-tools/swg/swagger"
	"gopkg.in/yaml.v2"
)

const (
	OK = "ok"
)

// DocGenarator 文档生成器
type DocGenarator struct {
	Conf             *Config
	Swagger          *swagger.Swagger
	globalParameters map[string]Parameter
	errs             []error
}

// NewDocGenarator 新建一个生成器
func NewDocGenarator(conf *Config) *DocGenarator {
	e := &DocGenarator{
		Conf:    conf,
		Swagger: &swagger.Swagger{},
	}
	e.initSwaggerConf()
	return e
}

// GET is a shortcut for gin router.Handle("GET", path, handle).
func (e *DocGenarator) GET(relativePath string, doc APIDoc) {
	// handle
	err := e.handle(GET, relativePath, doc)
	if err != nil {
		e.appendErr(err)
	}
}

// POST is a shortcut for gin router.Handle("POST", path, handle).
func (e *DocGenarator) POST(relativePath string, doc APIDoc) {
	// handle
	err := e.handle(POST, relativePath, doc)
	if err != nil {
		e.appendErr(err)
	}
}

// PUT is a shortcut for gin router.Handle("PUT", path, handle).
func (e *DocGenarator) PUT(relativePath string, doc APIDoc) {
	err := e.handle(PUT, relativePath, doc)
	if err != nil {
		e.appendErr(err)
	}
}

// PATCH is a shortcut for gin router.Handle("PATCH", path, handle).
func (e *DocGenarator) PATCH(relativePath string, doc APIDoc) {
	err := e.handle(PATCH, relativePath, doc)
	if err != nil {
		e.appendErr(err)
	}
}

// DELETE is a shortcut for gin router.Handle("DELETE", path, handle).
func (e *DocGenarator) DELETE(relativePath string, doc APIDoc) {
	err := e.handle(DELETE, relativePath, doc)
	if err != nil {
		e.appendErr(err)
	}
}

// GetSwaggerJSONDocument get swagger JSON document
func (e DocGenarator) GetSwaggerJSONDocument() (string, error) {
	data, err := json.MarshalIndent(e.Swagger, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// GetSwaggerYAMLDocument get  swagger YAML document
func (e DocGenarator) GetSwaggerYAMLDocument() (string, error) {
	data, err := yaml.Marshal(e.Swagger)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (e DocGenarator) GetSwaggerMarkdown() string {
	return e.Swagger.ToMarkdown()
}

func (e *DocGenarator) initSwaggerConf() {
	// set swagger Version
	e.Swagger.SwaggerVersion = "2.0"
	// set swagger Info
	e.Swagger.Info = &swagger.Info{
		Title:       e.Conf.Title,
		Description: e.Conf.Description,
		Version:     e.Conf.Version,
	}

	// set Host
	e.Swagger.Host = e.Conf.Host

	// set swagger basePath
	e.Swagger.BasePath = e.Conf.BasePath

	e.Swagger.Schemes = []string{}
	for _, scheme := range e.Conf.Schemes {
		e.Swagger.Schemes = append(e.Swagger.Schemes, string(scheme))
	}
}

func (e *DocGenarator) setSwaggerPath(relativePath string, method string, doc APIDoc) error {
	// to swagger path
	relativePath, err := ginPathToSwaggerPath(relativePath)
	if err != nil {
		return &engineError{relativePath, method, err}
	}
	// set swagger Paths
	operation, err := doc.ToSwaggerOperation()
	if err != nil {
		return &engineError{relativePath, method, err}
	}
	e.setSwaggerOperation(relativePath, method, operation)

	parameters, err := e.getParamters(operation.Parameters)
	if err != nil {
		return &engineError{relativePath, method, err}
	}

	// check paramter in relativePath
	if err := checkParametersInPath(relativePath, parameters); err != nil {
		return &engineError{relativePath, method, err}
	}

	// set swagger Definitions
	definitions, err := doc.ToSwaggerDefinitions()
	if err != nil {
		return &engineError{relativePath, method, err}
	}
	e.setSwaggerDefinitions(definitions)
	return nil
}

func (e *DocGenarator) getParamters(srcParameters []*swagger.Parameter) ([]*swagger.Parameter, error) {
	parameters := []*swagger.Parameter{}
	for _, parameter := range srcParameters {
		if parameter.Ref == "" {
			parameters = append(parameters, parameter)
			continue
		}
		if strings.HasPrefix(parameter.Ref, "#/parameters/") {
			name := parameter.Ref[len("#/parameters/"):]
			p, ok := e.globalParameters[name]
			if !ok {
				return nil, errors.New("the ref " + parameter.Ref + " is not found")
			}
			parameter, err := p.ToSwaggerParameters(name)
			if err != nil {
				return nil, err
			}
			parameters = append(parameters, parameter[0])
		}
	}
	return parameters, nil
}

func (e *DocGenarator) setSwaggerDefinitions(definitions map[string]*swagger.Schema) {
	for k, v := range definitions {
		// init swagger Definitions
		if e.Swagger.Definitions == nil {
			e.Swagger.Definitions = map[string]*swagger.Schema{}
		}
		// set value
		e.Swagger.Definitions[k] = v
	}
}

func (e *DocGenarator) setSwaggerOperation(relativePath string, method string, operation *swagger.Operation) {
	// init swagger paths
	if e.Swagger.Paths == nil {
		e.Swagger.Paths = make(map[string]*swagger.Item, 0)
	}
	// find swagger item
	item, ok := e.Swagger.Paths[relativePath]
	if !ok {
		item = &swagger.Item{}
		e.Swagger.Paths[relativePath] = item
	}
	// set operation
	switch method {
	case GET:
		item.Get = operation
	case POST:
		item.Post = operation
	case PUT:
		item.Put = operation
	case PATCH:
		item.Patch = operation
	case DELETE:
		item.Delete = operation
	}
}

func (e *DocGenarator) appendErr(err error) {
	if e.errs == nil {
		e.errs = []error{}
	}
	e.errs = append(e.errs, err)
}

// 打印错误
func (e *DocGenarator) PrintErrs() {
	errsDesc := e.GetErrsDesc()
	if errsDesc == "" {
		errsDesc = OK
	}
	fmt.Println(errsDesc)
}

func (e *DocGenarator) GetErrsDesc() string {
	len := len(e.errs)
	if len == 0 {
		return ""
	}
	errsDesc := fmt.Sprintf("发现%d个问题：\n", len)
	for i, err := range e.errs {
		errsDesc += fmt.Sprintf("  [%d] %s\n", i+1, err.Error())
	}
	return errsDesc
}

type engineError struct {
	relativePath string
	method       string
	err          error
}

func (e engineError) Error() string {
	return e.method + " " + e.relativePath + " " + e.err.Error()
}

func (e *DocGenarator) getBasePath() string {
	return e.Conf.BasePath
}

func (e *DocGenarator) SetGlobalParameters(parameters map[string]Parameter) error {
	e.globalParameters = parameters
	if len(parameters) > 0 {
		e.Swagger.Parameters = map[string]*swagger.Parameter{}
	}
	for name, paramter := range parameters {
		p, err := paramter.ToSwaggerParameters(name)
		if err != nil {
			return err
		}
		if len(p) != 1 {
			return errors.New("Invalid Global Parameter " + name)
		}
		e.Swagger.Parameters[name] = p[0]
	}
	return nil
}

func (e *DocGenarator) handle(method string, relativePath string, doc APIDoc) error {
	if doc == nil {
		return &engineError{relativePath, GET, errors.New("miss doc")}
	}

	// set method
	doc.SetMethod(method)

	// set swagger paths
	if err := e.setSwaggerPath(relativePath, method, doc); err != nil {
		return err
	}

	return nil
}
