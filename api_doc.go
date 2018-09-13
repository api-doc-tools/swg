package swg

import (
	"errors"
	"fmt"
	"sort"

	"github.com/api-doc-tools/swg/swagger"
)

// APIDoc 生成swagger文档的接口
type APIDoc interface {
	ToSwaggerOperation() (*swagger.Operation, error)
	ToSwaggerDefinitions() (map[string]*swagger.Schema, error)
	GetParameters() map[string]Parameter
	SetMethod(string)
}

// APIDocCommon 通用api文档的一个实现
// Fields:
//   Tags -- swagger tag分类.
//   Summary -- api简述
//   Description -- api详细的描述信息
//   Produces -- A list of MIME types the operation can produce.
//   Consumes -- A list of MIME types the operation can consume.
//   GlobalParameterNames --  swagger 文档里的全局参数
//   Parameters -- api下的个性参数，我们用不上，忽略， 使用GlobalParameterNames就可以了
//   Request -- 请求参数.
//   Responses -- 响应参数列表
type APIDocCommon struct {
	Tags                 []string
	Summary              string
	Description          string
	Produces             []string
	Consumes             []string
	GlobalParameterNames []string
	Parameters           map[string]Parameter
	Request              *Request
	Responses            map[int]Response
	method               string
}

func (doc *APIDocCommon) SetMethod(method string) {
	doc.method = method
}

func (doc APIDocCommon) ToSwaggerOperation() (*swagger.Operation, error) {
	if err := doc.check(); err != nil {
		return nil, err
	}
	operation := &swagger.Operation{Summary: doc.Summary, Description: doc.Description}
	// set Tags
	if len(doc.Tags) > 0 {
		operation.Tags = doc.Tags
	}
	// set Produces
	if len(doc.Produces) > 0 {
		operation.Produces = doc.Produces
	}
	// set Consumes
	if len(doc.Consumes) > 0 {
		operation.Consumes = doc.Consumes
	}
	// init Parameters
	if len(doc.Parameters) > 0 || doc.Request != nil {
		operation.Parameters = []*swagger.Parameter{}
	}
	// set parameters
	names := doc.getParamterNamesSortByName()
	for _, name := range doc.GlobalParameterNames {
		operation.Parameters = append(operation.Parameters, &swagger.Parameter{
			Ref: "#/parameters/" + name,
		})
	}
	for _, name := range names {
		parameters, err := doc.Parameters[name].ToSwaggerParameters(name)
		if err != nil {
			return nil, &parameterError{name, err}
		}
		operation.Parameters = append(operation.Parameters, parameters...)
	}

	if doc.Request != nil {
		param, err := doc.Request.toSwaggerParameter()
		if err != nil {
			return nil, err
		}
		operation.Parameters = append(operation.Parameters, param)
	}
	// init Responses
	if len(doc.Responses) > 0 {
		operation.Responses = map[string]*swagger.Response{}
	}
	// set Responses
	for statusCode, response := range doc.Responses {
		code := fmt.Sprintf("%d", statusCode)
		if statusCode == -1 {
			code = "default"
		}
		swaggerResponse, err := response.ToSwaggerResponse()
		if err != nil {
			return nil, err
		}
		operation.Responses[code] = swaggerResponse
	}
	return operation, nil
}

func (doc APIDocCommon) ToSwaggerDefinitions() (map[string]*swagger.Schema, error) {
	creater := StructDocCreater{}
	structDocs := map[string]*StructDoc{}
	// Request
	if doc.Request != nil && doc.Request.Model != nil {
		_structDocs, err := creater.GetStructDocMap(doc.Request.Model)
		if err != nil {
			return nil, err
		}
		for k, v := range _structDocs {
			structDocs[k] = v
		}
	}
	// range Responses
	for _, response := range doc.Responses {
		if response.Model != nil {
			_structDocs, err := creater.GetStructDocMap(response.Model)
			if err != nil {
				return nil, err
			}
			for k, v := range _structDocs {
				structDocs[k] = v
			}
		}
	}
	return getDefinitionsFromStructDocMap(structDocs), nil
}

func (doc APIDocCommon) GetParameters() map[string]Parameter {
	return doc.Parameters
}

func (doc APIDocCommon) check() error {
	if doc.hasformData() {
		if doc.Request != nil {
			return errors.New("There are parameters in formData, doc.Request should be nil")
		}
		if doc.method == GET {
			return errors.New("In method GET, param.InFormData should be nil")
		}
	}
	if doc.Request != nil {
		if doc.Request.Model == nil {
			return errors.New("doc.Request should not be nil")
		}
		if doc.method == GET {
			return errors.New("In method GET, doc.Request should be nil")
		}
	}
	return nil
}

func (doc APIDocCommon) hasformData() bool {
	for _, param := range doc.Parameters {
		if param.InFormData != nil {
			return true
		}
	}
	return false
}

func (doc APIDocCommon) getParamterNamesSortByName() []string {
	names := []string{}
	for name := range doc.Parameters {
		names = append(names, name)
	}
	sort.Sort(sort.StringSlice(names))
	return names
}
