package swagger

// 还在写

import (
	"fmt"
	"sort"
)

type FieldDoc struct {
	Name    string `desc:"参数名"`
	Type    string `desc:"值类型"`
	Range   string `desc:"取值范围"`
	Default string `desc:"默认值"`
	Example string `desc:"取值例子"`
	Desc    string `desc:"描述"`
}

func (field FieldDoc) ToTableLine() string {
	return fmt.Sprintf("| %s | %s | %s | %s | %s | %s |\n", field.Name, field.Type, field.Range, field.Default, field.Example, field.Desc)
}

type ModelDoc struct {
	Name      string     `desc:"model名称"`
	FieldDocs []FieldDoc `desc:"字段"`
}

func (model ModelDoc) ToModelTable() string {
	doc := "##" + model.Name + "\n"
	doc += "| 参数名 | 值类型 | 取值范围 | 默认值 | 取值例子 | 说明 |\n"
	doc += "| ----- | ----- | ------- | ----- | ------ | ---- |\n"
	for _, fieldDoc := range model.FieldDocs {
		doc += fieldDoc.ToTableLine()
	}
	return doc
}

func (swg Swagger) ToOnesMarkdown() string {
	modelNames := []string{}
	for k := range swg.Definitions {
		modelNames = append(modelNames, k)
	}
	sort.Sort(sort.StringSlice(modelNames))
	fmt.Println(modelNames)

	modelDocs := []ModelDoc{}
	for _, modelName := range modelNames {
		modelDocs = append(modelDocs, ModelDoc{
			Name:      modelName,
			FieldDocs: swg.ToFieldDocs(modelName),
		})
	}
	fmt.Println(modelDocs)
	doc := ""
	for _, modelDoc := range modelDocs {
		doc += modelDoc.ToModelTable()
	}
	return doc
}

func (swg Swagger) ToFieldDocs(modelName string) []FieldDoc {
	definition := swg.Definitions[modelName]
	fieldDocs := []FieldDoc{}
	fieldNames := []string{}
	for name := range definition.Properties {
		fieldNames = append(fieldNames, name)
	}
	sort.Sort(sort.StringSlice(fieldNames))
	for _, name := range fieldNames {
		field := definition.Properties[name]
		fieldDoc := field.ToFieldDoc(name)
		fieldDocs = append(fieldDocs, fieldDoc)
	}
	return fieldDocs
}

func (pro *Propertie) ToFieldDoc(name string) FieldDoc {
	Type := pro.Type
	if pro.Format != "" {
		Type += fmt.Sprintf("(%s)", pro.Format)
	}
	Default := ""
	if pro.Default != nil {
		Default = fmt.Sprintf("%v", pro.Default)
	}
	desc := pro.Description
	if pro.Ref != "" {
		desc += pro.Ref
	}

	return FieldDoc{
		Name:    name,
		Type:    Type,
		Range:   pro.getRange(),
		Default: Default,
		Example: pro.Example,
		Desc:    desc,
	}
}

func (pro Propertie) getRange() string {
	Range := ""
	if pro.Enum != nil {
		for i, enum := range pro.Enum {
			if i != 0 {
				Range += ","
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
	if pro.Minimum != nil {
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
