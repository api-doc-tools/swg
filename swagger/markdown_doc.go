package swagger

// 还在写

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type FieldDoc struct {
	Name     string `desc:"参数名"`
	Required string `desc:"是否必填"`
	Type     string `desc:"类型"`
	Range    string `desc:"取值范围"`
	Default  string `desc:"默认值"`
	Example  string `desc:"取值例子"`
	Desc     string `desc:"描述"`
}

func (field FieldDoc) ToTableLine() string {
	return fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n",
		field.Name, field.Required, field.Type, field.Range, field.Default, field.Example, field.Desc)
}

type ModelDoc struct {
	Name      string     `desc:"model名称"`
	FieldDocs []FieldDoc `desc:"字段"`
}

func (model ModelDoc) ToModelTable() string {
	doc := "## " + model.Name + "\n"
	doc += "| 参数名 | 必填 | 类型 | 取值范围 | 默认值 | 取值例子 | 说明 |\n"
	doc += "| ----- | ---- | ---- | ----- | ------- | ----- | ------ |\n"
	for _, fieldDoc := range model.FieldDocs {
		doc += fieldDoc.ToTableLine()
	}
	return doc
}

func (swg Swagger) ToMarkdown() string {
	doc := swg.getInfoDoc()
	doc += swg.getApisDoc()
	doc += swg.getModelsDoc()
	return doc
}

func (swg Swagger) getApisDoc() string {
	doc := "# API 列表\n\n"
	paths := []string{}
	for path := range swg.Paths {
		paths = append(paths, path)
	}
	sort.Sort(sort.StringSlice(paths))
	for _, path := range paths {
		item := swg.Paths[path]
		doc += swg.itemToDoc(path, item)
	}
	return doc
}

func (swg Swagger) itemToDoc(path string, item *Item) string {
	doc := ""
	if item.Get != nil {
		doc += swg.operationToDoc("GET", path, item.Get)
	}
	if item.Post != nil {
		doc += swg.operationToDoc("POST", path, item.Post)
	}
	if item.Delete != nil {
		doc += swg.operationToDoc("DELETE", path, item.Delete)
	}
	if item.Patch != nil {
		doc += swg.operationToDoc("PATCH", path, item.Patch)
	}
	if item.Put != nil {
		doc += swg.operationToDoc("PUT", path, item.Put)
	}
	return doc
}

func (swg Swagger) operationToDoc(method, path string, op *Operation) string {
	doc := fmt.Sprintf("## %s\n\n", op.Summary)
	if op.Description != "" {
		doc += op.Description + "\n"
	}
	doc += fmt.Sprintf("### URL\n")
	doc += fmt.Sprintf("%s %s\n", method, path)
	if len(op.Produces) > 0 {
		doc += fmt.Sprintf("### 请求方式\n")
		doc += fmt.Sprintf("%s\n", stringstoString(op.Produces))
	}
	if len(op.Consumes) > 0 {
		doc += fmt.Sprintf("### 响应方式\n")
		doc += fmt.Sprintf(" %s\n", stringstoString(op.Consumes))
	}
	doc += swg.parametersToDoc(op.Parameters)
	doc += swg.responsesToDoc(op.Responses)
	return doc
}

func (swg Swagger) responsesToDoc(responses map[string]*Response) string {
	doc := "### 响应参数列表\n"
	keys := []string{}
	for key := range responses {
		keys = append(keys, key)
	}
	sort.Sort(sort.StringSlice(keys))
	for _, key := range keys {
		resp := responses[key]
		doc += fmt.Sprintf("\n * %s:", key)
		if resp.Schema != nil {
			doc += fmt.Sprintf(" %s\n", toRefDoc(resp.Schema.Ref))
		} else {
			doc += fmt.Sprintf(" 无\n")
		}
		doc += fmt.Sprintf("%s\n", resp.Description)
	}
	doc += "\n"
	return doc
}

func (swg Swagger) parametersToDoc(parameters []*Parameter) string {
	doc := "### 参数列表\n"
	doc += "| 参数名 |  IN   | 必填 | 类型 | 取值范围 | 默认值 | 取值例子 | 说明 |\n"
	doc += "| ----- | ----- |---- | ---- | ----- | ------- | ----- | ------ |\n"
	var body *Parameter = nil
	for _, parmaeter := range parameters {
		if parmaeter.Ref != "" {
			doc += swg.parameterRefToDoc(parmaeter.Ref)
		} else {
			doc += parmaeter.ToTableLine()
			if parmaeter.Name == "body" {
				body = parmaeter
			}
		}
	}
	if body != nil {
		doc += "#### body说明\n"
		doc += body.Description
	}
	doc += "\n"
	return doc
}

func (swg Swagger) parameterRefToDoc(ref string) string {
	parameterName := strings.Replace(ref, "#/parameters/", "", -1)
	for name, parameter := range swg.Parameters {
		if name == parameterName {
			return parameter.ToTableLine()
		}
	}
	return ""
}

func (swg Swagger) getInfoDoc() string {
	doc := "# " + swg.Info.Title + "\n\n"
	doc += fmt.Sprintf("API版本: %s\n", swg.Info.Version)
	doc += fmt.Sprintf("Schemes: %s\n", stringstoString(swg.Schemes))
	doc += fmt.Sprintf("Host: %s\n", swg.Host)
	doc += fmt.Sprintf("BasePath: %s\n", swg.BasePath)
	doc += swg.Info.Description + "\n"
	return doc
}

func (swg Swagger) getModelsDoc() string {
	modelNames := []string{}
	for k := range swg.Definitions {
		modelNames = append(modelNames, k)
	}
	sort.Sort(sort.StringSlice(modelNames))

	modelDocs := []ModelDoc{}
	for _, modelName := range modelNames {
		modelDocs = append(modelDocs, ModelDoc{
			Name:      modelName,
			FieldDocs: swg.ToFieldDocs(modelName),
		})
	}
	doc := "# Model 列表\n"
	for _, modelDoc := range modelDocs {
		doc += modelDoc.ToModelTable()
	}
	return doc
}

func (swg Swagger) ToFieldDocs(modelName string) []FieldDoc {
	definition := swg.Definitions[modelName]
	fieldDocs := []FieldDoc{}
	fieldNames := []string{}
	requiredFieldNames := map[string]bool{}
	for _, name := range definition.Required {
		requiredFieldNames[name] = true
	}
	for name := range definition.Properties {
		fieldNames = append(fieldNames, name)
	}
	sort.Sort(sort.StringSlice(fieldNames))
	for _, name := range fieldNames {
		field := definition.Properties[name]
		fieldDoc := field.ToFieldDoc(name, requiredFieldNames[name])
		fieldDocs = append(fieldDocs, fieldDoc)
	}
	return fieldDocs
}

func (pro *Propertie) ToType() string {
	if pro.Ref != "" {
		return toRefDoc(pro.Ref)
	}
	if pro.Format != "" {
		return pro.Format
	}
	if pro.Type == "array" && pro.Items != nil {
		Type := "array "
		if pro.Items.Ref != "" {
			return Type + toRefDoc(pro.Items.Ref)
		}
		if pro.Items.Format != "" {
			return Type + pro.Items.Format
		}
		return Type + pro.Items.Type
	}
	return pro.Type
}

func toRefDoc(ref string) string {
	if ref == "" {
		return ref
	}
	str := strings.Replace(ref, "#/definitions/", "", -1)
	return toGFMAnchor(str)
}

// Github Flavored Markdownn(GFM)
func toGFMAnchor(str string) string {
	lowerStr := strings.ToLower(str)
	spaceRegexp := regexp.MustCompile(`\s+`)
	anchor := spaceRegexp.ReplaceAllString(lowerStr, "-")
	return fmt.Sprintf("[%s](#%s)", str, anchor)
}

func (pro *Propertie) ToFieldDoc(name string, required bool) FieldDoc {
	defaultValue := ""
	if pro.Default != nil {
		defaultValue = fmt.Sprintf("%v", pro.Default)
	}

	return FieldDoc{
		Name:     name,
		Type:     pro.ToType(),
		Range:    pro.getRange(),
		Default:  defaultValue,
		Example:  pro.Example,
		Desc:     strings.Replace(pro.Description, "\n", "<br>", -1),
		Required: getRequiredDesc(required),
	}
}

func getRequiredDesc(required bool) string {
	if required {
		return "是"
	}
	return "否"
}

func (pro Propertie) getRange() string {
	Range := ""
	if pro.Enum != nil {
		for i, enum := range pro.Enum {
			if i != 0 {
				Range += "<br>"
			}
			Range += fmt.Sprintf("%v", enum)
		}
	}

	if pro.Minimum != nil {
		if Range != "" {
			Range += ", "
		}
		Range += fmt.Sprintf(">%f", *pro.Minimum)
	}
	if pro.Maximum != nil {
		if Range != "" {
			Range += ", "
		}
		Range += fmt.Sprintf(">%f", *pro.Maximum)
	}
	if pro.MinLength != nil && pro.MaxLength != nil {
		if *pro.MinLength == *pro.MaxLength {
			if Range != "" {
				Range += ", "
			}
			Range += fmt.Sprintf("len=%d", *pro.MinLength)
			return Range
		}
	}
	if pro.MinLength != nil {
		if Range != "" {
			Range += ", "
		}
		Range += fmt.Sprintf("len>=%d", *pro.MinLength)
	}
	if pro.MaxLength != nil {
		if Range != "" {
			Range += ", "
		}
		Range += fmt.Sprintf("len<=%d", *pro.MaxLength)
	}
	return Range
}

func (p Parameter) ToTableLine() string {
	if p.Name == "body" {
		return fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s | %s |\n",
			p.Name, p.In, getRequiredDesc(p.Required), toRefDoc(p.Schema.Ref), "", "", "", "")
	}
	defaultValue := ""
	if p.Default != nil {
		defaultValue = fmt.Sprintf("%v", p.Default)
	}
	//"| 参数名 |  IN   | 必填 | 类型 | 取值范围 | 默认值  | 取值例子 | 说明 |\n"
	return fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s | %s |\n",
		p.Name, p.In, getRequiredDesc(p.Required), p.Type, p.getRange(), defaultValue, "", strings.Replace(p.Description, "\n", "<br>", -1))
}

func (p Parameter) getRange() string {
	Range := ""
	if p.Enum != nil {
		for i, enum := range p.Enum {
			if i != 0 {
				Range += "<br>"
			}
			Range += fmt.Sprintf("%v", enum)
		}
	}

	if p.Minimum != nil {
		if Range != "" {
			Range += ", "
		}
		Range += fmt.Sprintf(">%f", *p.Minimum)
	}
	if p.Maximum != nil {
		if Range != "" {
			Range += ", "
		}
		Range += fmt.Sprintf(">%f", *p.Maximum)
	}
	if p.MinLength != nil && p.MaxLength != nil {
		if *p.MinLength == *p.MaxLength {
			if Range != "" {
				Range += ", "
			}
			Range += fmt.Sprintf("len=%d", *p.MinLength)
			return Range
		}
	}
	if p.MinLength != nil {
		if Range != "" {
			Range += ", "
		}
		Range += fmt.Sprintf("len>=%d", *p.MinLength)
	}
	if p.MaxLength != nil {
		if Range != "" {
			Range += ", "
		}
		Range += fmt.Sprintf("len<=%d", *p.MaxLength)
	}
	return Range
}

func stringstoString(strs []string) string {
	result := ""
	for i, str := range strs {
		if i != 0 {
			result += ","
		}
		result += str
	}
	return result
}
