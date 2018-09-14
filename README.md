# api 文档工具

通过go代码来生成swagger文档

假如文档生成有问题， 可以通过 DocGenarator.PrintlnErrs() 打印收集到的错误，错误提示的举例：
```
发现1个问题：
  [1] GET /book/{bookUUID}/{uuid} miss parameters [uuid] in APIDoc
```

本人很懒，文档回头补正确

可以先看这个例子：
https://github.com/api-doc-tools/swg/tree/master/examples/apis

- [api 文档工具](#api-文档工具)
    - [使用说明](#使用说明)
            - [Model 的字段](#model-的字段)
                - [字段名](#字段名)
                - [字段类型](#字段类型)
                - [字段标签](#字段标签)
                    - [字段标签-json,xml,dec](#字段标签-jsonxmldec)
                    - [字段标签-枚举(enum)](#字段标签-枚举enum)
                    - [字段标签-最小值（min）、最大值（max）](#字段标签-最小值min最大值max)
                - [字段标签-最小长度（minlen）、最大长度（maxlen）](#字段标签-最小长度minlen最大长度maxlen)
        - [APIDoc](#apidoc)
            - [APIDoc参数demo（不要用这个了， 我们有全局参数的方法会更加方便）](#apidoc参数demo不要用这个了-我们有全局参数的方法会更加方便)
            - [APIDoc Demo](#apidoc-demo)
        - [Config](#config)
            - [Config Demo](#config-demo)
    - [API Demo](#api-demo)
        - [Upload files](#upload-files)
        - [POST,PUT,PATCH,DELETE](#postputpatchdelete)
        - [GET](#get)
            - [GET - JSON](#get---json)
            - [GET - XML](#get---xml)
            - [GET - JSON And XML](#get---json-and-xml)
            - [GET - download file](#get---download-file)


## 使用说明

#### Model 的字段

一个字段由三部分组成： [字段名] [字段的类型] [字段的标签]
例如 
```golang
type Book struct {
	ID        string        `json:"id" desc:"书的id"`
	Images    BookImageUrls `json:"images" desc:"书的图片"`
}
```
Model Book 有两个字段，分别是： ID 和 Images. 
字段`ID`的类型是`string`，标签是：`json:"id" desc:"书的id"`
字段`Images`的类型是`BookImageUrls`,标签是:`json:"images" desc:"书的图片"`

##### 字段名

字段名要以大写字母开头。

##### 字段类型

仅支持以下几种 ：
``` golang
  int, int32, int64, uint, uint32, uint64, bool, string, float32, float64,
  *int, *int32, *int64, *uint, *uint32, *uint64, *bool, *string,  *float32, *float64,
  []int, []int32, []int64, []uint, []uint32, []uint64, []bool, []string,  []float32, []float64,
  []*int, []*int32, []*int64, []*uint, []*uint32, []*uint64, []*bool, []*string,  []*float32, []*float64,
  struct, []struct, []*struct 
```
##### 字段标签

目前以支持的标签有:  描述desc、最小值min、最大值max、最小长度minlen、最大长度maxlen、默认值default
计划准备添加特性： 是否必填req、正则Pattern

例子: 
```golang
type XXXDemo struct {
	ID      string   `json:"id" xml:"id" desc:"the id"`
	score   float32  `json:"price" xml:"price" min:"0" max:"99.9" desc:"the score"`
	Type    string   `json:"type" xml:"type" enum:"type1 typ2"`
	Valid   bool     `json:"valid" xml:"valid" desc:"is valid"`
	Authors []string `json:"authors" desc:"the authors"`
}

type YYYDemo struct {
	Offset int64      `json:"offset" xml:"offset" desc:"the offset"`
	Limit  int64      `json:"limit" xml:"limt" enum:"0, 10, 100, 1000" desc:"the limit"`
	XXXs   []*XXXDemo `json:"xxxs" xml:"xxxs" desc:"XXXs"`
}
```

标签有以下几个： `json,xml,desc,enum,max,min`， 即 json的键名、xml的键名， 该字段的描述信息、枚举、最大值、最小值。
每一个标签都是可选的，属性之间用空格分隔。

###### 字段标签-json,xml,dec

- json: json的键名.
- xml: xml的键名.
- desc: 该字段的描述信息.

###### 字段标签-枚举(enum)
枚举的例子:
```golang
type User struct {
	Name string
	Type string `enum:"admin normal"`
}
```
枚举值之间用空格分隔。
只有以下类型是支持枚举的： `int, int32, int64, uint, uint32, uint64, string`

(框架会自动检查给出的枚举值跟字段类型是否匹配 ，如何不匹配，则会指出出错发生的位置)

###### 字段标签-最小值（min）、最大值（max）

只有类型是数字（`int, int32, int64, uint, uint32, uint64, float32, float64`）时才允许设置最大值和最小值标签

#####  字段标签-最小长度（minlen）、最大长度（maxlen）

只有string才能用这个标签

### APIDoc


#### APIDoc参数demo（不要用这个了， 我们有全局参数的方法会更加方便）

```golang
swg.APIDocCommon{
	Parameters: map[string]swg.Parameter{
		// 参数名称是id，参数出现在http请求的 Path 中，类型是string， 描述是"the id"
		// 例如: 在 /items/{id}, 参数id出现在 http请求的Path中
		"id": swg.Parameter{InPath: &swg.Value{Type: "string", Desc: "the id"}},
		// 参数名称是user-type，参数出现在http请求的header中，类型是string， 枚举类型有: admirn 和 normal, 该参数是必填的， 描述是"user type"
		"user-type": swg.Parameter{InHeader: &swg.Value{Type: "string", Enum: "admin normal", Required: true, Desc: "user type"}},
		// 参数名是limit, 参数出现在http请求的Query中， 类型是int32, 最小值是0, 最大值是1000, 该参数是必填的
		// 例如： 在 /items?limit=###, 参数limit 出现在http请求的Query中
		"limit": swg.Parameter{InQuery: &swg.Value{Type: "int32", Min: "0", Max: "100", Required: true}},
		// 参数名是data, 参数出现在http请求的FormData中， 类型是string
		"data":  swg.Parameter{InFormData: &swg.Value{Type: "string"},
		// 参数名是file1, 参数出现在http请求的formData中， 类型是file, 最小值是0, 最大值是1000, 该参数是选填的
	    "file1": swg.Parameter{InFormData: &swg.Value{Type: "file", Desc: "the file to upload"}},
	},
}
```
APIDoc参数在http请求中的位置有:`InPath,InHeader,InQuery,InFormData`.

Value.Type支持以下类型:`int, int32, int64, uint, uint32, uint64, bool, string, float32, float64, file`.

注意： 在非InFormData的情况下，Value.Type不允许设置为`file `


####  APIDoc Demo

用一个APIDoc来描述一个API
```golang

// Model ErrorMessage
type ErrorMessage struct {
	Message string `json:"message" desc:"the error message"`
	Details string `json:"detail" desc:"the error detail"`
}

// Model Demo
type Demo struct {
	ID    string `json:"id" desc:"the id"`
	Title string `json:"title" desc:"the title"`
}

var DocCommonDemo = &swg.APIDocCommon{
	// 选填， 该API的标签，仅用于api文档的分类显示
	Tags: []string{"demo"},
	// 选填，该API的简述
	Summary: "a demo api summary",
	// 选填： 该API可以产生的MIME类型列表
	Produces: []string{swg.Application_Json},
	// 选填: 该API可以消费的MIME类型列表, (注意，如果method 是 GET,Consumes不可填)
	Consumes: []string{swg.Application_Json},
	// 选填: http 请求中的参数列表, 参数的类型支持：int, int32, int64, uint, uint32, uint64, bool, string, float32, float64, file
	// (只有InFormData时,file类型才被允许，其他情况不允许类型为file)
	Parameters: map[string]swg.Parameter{
		// 参数名称是id，参数出现在http请求的 Path 中，类型是string， 描述是"the id"
		// 例如: 在 /items/{id}, 参数id出现在 http请求的Path中
		"id": swg.Parameter{InPath: &swg.Value{Type: "string", Desc: "the id"}},
		// 参数名称是user-type，参数出现在http请求的header中，类型是string， 枚举类型有: admirn 和 normal, 该参数是必填的， 描述是"user type"
		"user-type": swg.Parameter{InHeader: &swg.Value{Type: "string", Enum: "admin normal", Required: true, Desc: "user type"}},
		// 参数名是limit, 参数出现在http请求的Query中， 类型是int32, 最小值是0, 最大值是1000, 该参数是必填的
		// 例如： 在 /items?limit=###, 参数limit 出现在http请求的Query中
		"limit": swg.Parameter{InQuery: &swg.Value{Type: "int32", Min: "0", Max: "100", Required: true}},
		// 参数名是data, 参数出现在http请求的FormData中， 类型是string
		// "data":  swg.Parameter{InFormData: &swg.Value{Type: "string"},
		// 参数名是file1, 参数出现在http请求的formData中， 类型是file, 最小值是0, 最大值是1000, 该参数是选填的
	    // "file1": swg.Parameter{InFormData: &swg.Value{Type: "file", Desc: "the file to upload"}},
	},
	// 随请求一起发送的Model (注意，1.如果method 是 GET,Request不可填；2.如果Parameters存在InFormData的参数，Request也不可填)
	Request: &swg.Request{
		Description: "the demo info",
		Model:       &Demo{},
	},
	// http 响应列表
	Responses: map[int]swg.Response{
		// StatusCode 是 200 时
		200: swg.Response{
			// 该响应的描述
			Description: "successful operation",
			//随响应一起发送的Model
			Model: &Demo{},
		},
		// StatusCode 是 400 时
		400: swg.Response{
			// 该响应的描述
			Description: "failed operation",
			//随响应一起发送的Model
			Model: &ErrorMessage{},
			//随响应一起发送的标头列表。
			Headers: map[string]swg.Value{
				"xxx": swg.Value{Type: "string"},
			},
		},
	},
}

```

### Config

#### Config Demo

```golang
conf := &swg.Config{
	// 必填，API的传输协议。 从列表swg.Schemswg, swg.SchemswgS中选择
	Schemes:            []swg.Scheme{swg.Schemswg, swg.SchemswgS},
	// 选填，API的基本路径， 必须斜杠(/)开头，BasePath不支持路径模板。
	BasePath:           "/dev",
	// 必填: 版本号（目前在框架中，仅用于文档显示，未用于逻辑）
	Version:            "v1",
	// 必填: APIs文档的标题
	Title:              "Demo APIS",
	// 必填: APIs文档的描述信息
    Description:        "Demo APIS Description",
    // 必填: 域名或ip
    DomainName:         "xxx.com",
}
```

## API Demo

### Upload files

```go

var doc = &swg.APIDocCommon{
	Summary:  "doc summary",
	Produces: []string{swg.Application_Json},
	Consumes: []string{swg.Application_Json},
	Parameters: map[string]swg.Parameter{
		"file": swg.Parameter{InFormData: &swg.Value{Type: "file", Desc: "the file to upload"}},
	},
	Responses: map[int]swg.Response{
		200: swg.Response{
			Description: "successful operation",
		},
		400: swg.Response{
			Description: "failed operation",
		},
	},
}

```

### POST,PUT,PATCH,DELETE

```golang
type ErrorMessage struct {
	Message string
}

type XXXReq struct {
	ID string
}

type XXXRsp struct {
	ID   string
	Name string
}

var doc = &swg.APIDocCommon{
	Summary:  "doc summary",
	Produces: []string{swg.Application_Json},
	Consumes: []string{swg.Application_Json},
	Parameters: map[string]swg.Parameter{
		"version": swg.Parameter{InHeader: &swg.Value{Type: "string", Desc: "the version of api"}},
	},
	Request: &swg.Request{
		Description: "request model",
		Model:       &XXXReq{},
	},
	Responses: map[int]swg.Response{
		200: swg.Response{
			Description: "successful operation",
			Model:       &XXXRsp{},
		},
		400: swg.Response{
			Description: "failed operation",
			Model:       &ErrorMessage{},
		},
	},
}

```

### GET

#### GET - JSON

```golang

type ErrorMessage struct {
	Message string `json:"message" desc:"the error message"`
	Details string `json:"detail" desc:"the error detail"`
}

type Book struct {
	ID    string `json:"id" desc:"the book id"`
	Title string `json:"title" desc:"the book title"`
}

var DocGETBook = &swg.APIDocCommon{
	Summary:  "Get book info by id",
	Produces: []string{swg.Application_Json},
	Parameters: map[string]swg.Parameter{
		"id": swg.Parameter{InPath: &swg.Value{Type: "string", Desc: "the id of book"}},
	},
	Responses: map[int]swg.Response{
		200: swg.Response{
			Description: "successful operation",
			Model:       &Book{},
		},
		400: swg.Response{
			Description: "failed operation",
			Model:       &ErrorMessage{},
		},
	},
}

```

#### GET - XML
```golang

type ErrorMessage struct {
	Message string `xml:"message" desc:"the error message"`
	Details string `xml:"detail" desc:"the error detail"`
}

type Book struct {
	ID    string `xml:"id" desc:"the book id"`
	Title string `xml:"title" desc:"the book title"`
}

var DocGETBook = &swg.APIDocCommon{
	Summary:  "Get book info by id",
	Produces: []string{swg.Application_Xml},
	Parameters: map[string]swg.Parameter{
		"id": swg.Parameter{InPath: &swg.Value{Type: "string", Desc: "the id of book"}},
	},
	Responses: map[int]swg.Response{
		200: swg.Response{
			Description: "successful operation",
			Model:       &Book{},
		},
		400: swg.Response{
			Description: "failed operation",
			Model:       &ErrorMessage{},
		},
	},
}

```

#### GET - JSON And XML

```golang

type ErrorMessage struct {
	Message string `json:"message" xml:"message" desc:"the error message"`
	Details string `json:"detail" xml:"detail" desc:"the error detail"`
}

type Book struct {
	ID    string `json:"id" xml:"id" desc:"the book id"`
	Title string `json:"title" xml:"title" desc:"the book title"`
}

var DocGETBook = &swg.APIDocCommon{
	Summary:  "Get book info by id",
	Produces: []string{swg.Application_Json, swg.Application_Xml},
	Parameters: map[string]swg.Parameter{
		"id": swg.Parameter{InPath: &swg.Value{Type: "string", Desc: "the id of book"}},
	},
	Responses: map[int]swg.Response{
		200: swg.Response{
			Description: "successful operation",
			Model:       &Book{},
		},
		400: swg.Response{
			Description: "failed operation",
			Model:       &ErrorMessage{},
		},
	},
}

```

#### GET - download file

```golang

type ErrorMessage struct {
	Message string `json:"message" desc:"the error message"`
	Details string `json:"detail" desc:"the error detail"`
}

var DocDownloadText = &swg.APIDocCommon{
	Summary:  "A download file demo",
	Produces: []string{swg.Image_Jpeg},
	Parameters: map[string]swg.Parameter{
		"fileName": swg.Parameter{InPath: &swg.Value{Type: "string", Desc: "the fileName"}},
	},
	Responses: map[int]swg.Response{
		200: swg.Response{
			Description: "successful operation",
		},
		400: swg.Response{
			Description: "failed operation",
			Model:       &ErrorMessage{},
		},
	},
}

```