package main

import (
	"fmt"
	"io/ioutil"

	"github.com/api-doc-tools/swg"
)

// 目前以支持的标签有:  描述desc、最小值min、最大值max、最小长度minlen、最大长度maxlen、默认值default
// 新加的特性： 是否必填req、
// 计划准备添加特性正则Pattern
type APIError struct {
	Code    int    `json:"code" xml:"code" req:"true" desc:"错误编码"`
	Message string `json:"message" xml:"message" desc:"错误信息"`
}

type Image struct {
	ImagUUID   string `json:"image_uuid" xml:"image_uuid" desc:"图片uuid"`
	Size       int    `json:"size" xml:"size" max:"10240" desc:"图片大小"`
	CreateTime int64  `json:"create_time" xml:"create_time" desc:"创建时间"`
}

type Book struct {
	UUID        string   `json:"id" xml:"id" maxlen:"8" desc:"书的uuid"`
	Type        string   `json:"type" xml:"type" enum:"scientific sociology physics" desc:"书的类型"`
	Title       string   `json:"title" xml:"title" desc:"书的题目"`
	Summary     string   `json:"summary" xml:"summary"  maxlen:"255" desc:"书的简介"`
	Authors     []string `json:"authors" xml:"authors" desc:"作者列表"`
	Images      []*Image `json:"images" xml:"images" desc:"图片urls"`
	Pages       int      `json:"pages" xml:"pages" desc:"页数"`
	Price       float32  `json:"price" xml:"price" min:"1.0" max:"9999.0" desc:"价格"`
	AlreadyRead bool     `json:"already_read,omitempty" xml:"already_read,omitempty" desc:"已读"`
}

type Books struct {
	Total int    `json:"total" xml:"total" desc:"总数"`
	Books []Book `json:"books" xml:"books" desc:""`
}

const (
	BookTypeScientific = "scientific"
	BookTypeSociology  = "sociology"
	BookTypePhysics    = "physics"
)

var BookTypeList = []string{
	BookTypeScientific,
	BookTypeSociology,
	BookTypePhysics,
}

var BookTypeMap = map[string]string{
	BookTypeScientific: "科学",
	BookTypeSociology:  "社会",
	BookTypePhysics:    "物理",
}

const (
	BookUUID = "bookUUID"
	UserUUID = "userUUID"
	Token    = "Token"
)

// 参数位置有4种： InPath InHeader InQuery InFormData
// 如果参数是InPath的，那么该参数是必须填的， 非InPath的参数，如果要求必填， 需要设置 Value.Required=true
// 支持设置 参数类型Type、枚举Enum、最小值Min、最大值Max、最小字符串长度MinLen、最大字符串长度MaxLen、描述信息Desc、默认值Default、是否必填Required、正则表达式Pattern
var GlobalParameters = map[string]swg.Parameter{
	BookUUID: swg.Parameter{InPath: &swg.Value{Type: "string", MaxLen: "8", Desc: "书的id"}},
	UserUUID: swg.Parameter{InHeader: &swg.Value{Type: "string", MaxLen: "8", Required: true, Desc: "用户uuid"}},
	Token:    swg.Parameter{InHeader: &swg.Value{Type: "string", MaxLen: "64", Required: true, Desc: "身份认证Token"}},
}

func BookTypeDesc() string {
	desc := "书的分类\n\n"
	desc += "| type | 描述 |\n"
	desc += "| ---- | ---- |\n"
	for _, bookType := range BookTypeList {
		desc += fmt.Sprintf("| %s | %s |\n", bookType, BookTypeMap[bookType])
	}
	return desc
}

// Get请求, Produces 不用设置, Request 不用设置
// 该例子中 Consumes 允许接口返回json或xml
// Responses 的key 为-1时， 对应swagger的default
var ListBooks = &swg.APIDocCommon{
	Tags:                 []string{"book"},
	Summary:              "列出所有的书",
	Consumes:             []string{swg.Application_Json, swg.Application_Xml},
	GlobalParameterNames: []string{UserUUID, Token},
	Responses: map[int]swg.Response{
		200: swg.Response{
			Description: "正确返回",
			Model:       &Books{},
		},
		-1: swg.Response{
			Description: "错误返回",
			Model:       &APIError{},
		},
	},
}

var GetBook = &swg.APIDocCommon{
	Tags:                 []string{"book"},
	Summary:              "获取一本书的信息",
	Consumes:             []string{swg.Application_Json, swg.Application_Xml},
	GlobalParameterNames: []string{UserUUID, Token, BookUUID},
	Responses: map[int]swg.Response{
		200: swg.Response{
			Description: "正确返回",
			Model:       &Book{},
		},
		-1: swg.Response{
			Description: "错误返回",
			Model:       &APIError{},
		},
	},
}

var AddBook = &swg.APIDocCommon{
	Tags:                 []string{"book"},
	Summary:              "上传书的信息",
	Consumes:             []string{swg.Application_Json, swg.Application_Xml},
	Produces:             []string{swg.Application_Json, swg.Application_Xml},
	GlobalParameterNames: []string{UserUUID, Token},
	Request: &swg.Request{
		Description: "正确结果",
		Model:       &Books{},
	},
	Responses: map[int]swg.Response{
		200: swg.Response{
			Description: "正确结果",
		},
		-1: swg.Response{
			Description: "错误返回",
			Model:       &APIError{},
		},
	},
}

func SaveSwaggerDoc() error {
	conf := &swg.Config{
		Schemes:     []swg.Scheme{swg.SchemeHTTP, swg.SchemeHTTPS}, // 该接可以用http或https访问
		BasePath:    "/store",
		Version:     "v1",
		Title:       "The Book Store Demo",
		Description: "这是一个例子， 展示了如何使用 swg包来生成swagger文档\n" + BookTypeDesc(),
		Host:        "api-doc-tools.com",
	}
	genarator := swg.NewDocGenarator(conf) // 新建一个生成器

	genarator.SetGlobalParameters(GlobalParameters) // 设置全局参数， 懒人必备

	genarator.GET("/books", ListBooks) // get post patch delete神马的都支持啊
	genarator.GET("/book/:bookUUID", GetBook)
	genarator.POST("/books", AddBook)

	genarator.PrintErrs() // 打印错误信息

	doc, err := genarator.GetSwaggerYAMLDocument() // swagger ymal格式文档
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("swagger.yaml", []byte(doc), 0644)
	if err != nil {
		return err
	}
	doc = genarator.GetSwaggerMarkdown() // 获取markdown格式文档
	err = ioutil.WriteFile("swagger.md", []byte(doc), 0644)
	if err != nil {
		return err
	}
	return nil
}
