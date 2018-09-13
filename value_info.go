package swg

import (
	"errors"
	"strconv"
	"strings"

	"github.com/api-doc-tools/swg/swagger"
)

// Value 参数信息
// Type -- (必填) 参数类型。仅支持以下类型:(int, int32, int64, uint,  uint32, uint64, float32, float64, string, bool)
// Enum -- 枚举。枚举类型之间通过空格分割. 仅支持以下类型: string, int, int32, int64, uint, uint32, uint64
// Min  -- 最小值。(仅支持以下类型: int, int32, int64, uint,  uint32, uint64, float32, float64)
// Max  -- 最大值。(仅支持以下类型: int, int32, int64, uint,  uint32, uint64, float32, float64)
// Desc -- 描述信息。
// MinLen -- string的最小长度
// MaxLen -- string的最大长度
// Defalut -- 默认值
// Required -- 该参数是否必填
// Pattern -- 正则表达式

type Value struct {
	Type     string
	Enum     string
	Min      string
	Max      string
	Desc     string
	MinLen   string
	MaxLen   string
	Default  string
	Required bool
	Pattern  string
}

func (v Value) checkWithHTTPIn(in string) error {
	if v.Type == "file" && in != InFormData {
		return errors.New("in HTTP " + in + ", value type can't be file")
	}
	return v.check()
}

func (v Value) check() error {
	if err := v.checkValuetype(); err != nil {
		return err
	}
	if v.hasEnum() {
		if err := v.checkEnum(); err != nil {
			return err
		}
	}
	if v.hasMin() {
		if err := v.checkMinimum(); err != nil {
			return err
		}
	}
	if v.hasMax() {
		if err := v.checkMaximum(); err != nil {
			return err
		}
	}
	if v.hasMax() && v.hasMin() {
		if err := compareMinimumAndMaximum(v.Min, v.Max, v.Type); err != nil {
			return err
		}
	}
	return nil
}

func (v Value) hasEnum() bool {
	return v.Enum != ""
}

func (v Value) hasMin() bool {
	return v.Min != ""
}

func (v Value) hasMax() bool {
	return v.Max != ""
}

func (v Value) checkValuetype() error {
	switch v.Type {
	case "string":
	case "int":
	case "int32":
	case "int64":
	case "uint":
	case "uint32":
	case "uint64":
	case "float32":
	case "float64":
	case "bool":
	case "file":
	default:
		return errors.New("the paramter value type is " + v.Type + ",it is not supported")
	}
	return nil
}

func (v Value) checkEnum() error {
	if v.Min != "" {
		return errors.New("Enum exists, cant't set Min")
	}
	if v.Max != "" {
		return errors.New("Enum exists, cant't set Max")
	}
	return checkEnumFormat(v.Enum, v.Type)
}

func (v Value) checkMinimum() error {
	if !v.isNumber() {
		return errors.New("the paramter value type is " + v.Type + ",  can't set Min and Min")
	}
	return checkMinimumFormat(v.Min, v.Type)
}

func (v Value) checkMaximum() error {
	if !v.isNumber() {
		return errors.New("the paramter value type is " + v.Type + ",  can't set Min and Min")
	}
	return checkMaximumFormat(v.Max, v.Type)
}

func (v Value) isNumber() bool {
	if v.isFloat() || v.isInt() || v.isUint() {
		return true
	}
	return false
}

func (v Value) isString() bool {
	switch v.Type {
	case "string":
	default:
		return false
	}
	return true
}

func (v Value) isBool() bool {
	switch v.Type {
	case "bool":
	default:
		return false
	}
	return true
}

func (v Value) isFloat() bool {
	switch v.Type {
	case "float32":
	case "float64":
	default:
		return false
	}
	return true
}

func (v Value) isInt() bool {
	switch v.Type {
	case "int":
	case "int32":
	case "int64":
	default:
		return false
	}
	return true
}

func (v Value) isUint() bool {
	switch v.Type {
	case "uint":
	case "uint32":
	case "uint64":
	default:
		return false
	}
	return true
}

func (v Value) getMinLen() (*int64, error) {
	if v.MinLen == "" {
		return nil, nil
	}
	num, err := strconv.ParseInt(v.MinLen, 10, 64)
	if err != nil {
		return nil, err
	}
	return &num, nil
}

func (v Value) getMaxLen() (*int64, error) {
	if v.MaxLen == "" {
		return nil, nil
	}
	num, err := strconv.ParseInt(v.MaxLen, 10, 64)
	if err != nil {
		return nil, err
	}
	return &num, nil
}

func (v Value) getEnum() ([]interface{}, error) {
	enum := []interface{}{}
	if v.Enum == "" {
		return enum, nil
	}
	if v.isInt() {
		_enum := strings.Fields(v.Enum)
		for _, v := range _enum {
			num, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return nil, err
			}
			enum = append(enum, num)
		}
	}
	if v.isUint() {
		_enum := strings.Fields(v.Enum)
		for _, v := range _enum {
			num, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				return nil, err
			}
			enum = append(enum, num)
		}
	}
	if v.isString() {
		_enum := strings.Fields(v.Enum)
		for _, v := range _enum {
			enum = append(enum, v)
		}
	}
	return enum, nil
}

func (v Value) getMinimum() (*float64, error) {
	if v.Min == "" {
		return nil, nil
	}
	if v.isFloat() || v.isInt() || v.isUint() {
		min, err := strconv.ParseFloat(v.Min, 64)
		if err != nil {
			return nil, err
		}
		return &min, err
	}
	return nil, nil
}

func (v Value) getMaximum() (*float64, error) {
	if v.Max == "" {
		return nil, nil
	}
	if v.isFloat() || v.isInt() || v.isUint() {
		min, err := strconv.ParseFloat(v.Max, 64)
		if err != nil {
			return nil, err
		}
		return &min, err
	}
	return nil, nil
}

func (v Value) toSwaggerHeader() (*swagger.Header, error) {
	if v.Type == "file" {
		return nil, errors.New("type file is not supported")
	}
	dataType, ok := dataTypes[v.Type]
	if !ok {
		return nil, errors.New("type " + v.Type + " is not supported")
	}
	min, err := v.getMinimum()
	if err != nil {
		return nil, err
	}
	max, err := v.getMaximum()
	if err != nil {
		return nil, err
	}
	enum, err := v.getEnum()
	if err != nil {
		return nil, err
	}
	if len(enum) == 0 {
		enum = nil
	}
	minLen, err := v.getMinLen()
	if err != nil {
		return nil, err
	}
	maxLen, err := v.getMaxLen()
	if err != nil {
		return nil, err
	}
	defalutValue, err := v.getDefalut()
	if err != nil {
		return nil, err
	}
	return &swagger.Header{
		Description: v.Desc,
		Type:        dataType.typeName,
		Format:      dataType.format,
		Minimum:     min,
		Maximum:     max,
		Enum:        enum,
		MinLength:   minLen,
		MaxLength:   maxLen,
		Default:     defalutValue,
	}, nil
}

func (v Value) getDefalut() (interface{}, error) {
	if v.Default == "" {
		return nil, nil
	}
	if v.isBool() {
		defaultValue, err := strconv.ParseBool(v.Default)
		if err != nil {
			return nil, err
		}
		return &defaultValue, nil
	}
	if v.isFloat() {
		defaultValue, err := strconv.ParseFloat(v.Default, v.getBitSize())
		if err != nil {
			return nil, err
		}
		return &defaultValue, nil
	}
	if v.isInt() {
		defaultValue, err := strconv.ParseInt(v.Default, 10, v.getBitSize())
		if err != nil {
			return nil, err
		}
		return &defaultValue, nil
	}
	if v.isString() {
		defaultValue := v.Default
		return &defaultValue, nil
	}
	if v.isUint() {
		defaultValue, err := strconv.ParseUint(v.Default, 10, v.getBitSize())
		if err != nil {
			return nil, err
		}
		return &defaultValue, nil
	}
	return nil, nil
}

func (v Value) getBitSize() int {
	switch v.Type {
	case "int":
		return 32
	case "int32":
		return 32
	case "int64":
		return 64
	case "uint":
		return 32
	case "uint32":
		return 32
	case "uint64":
		return 64
	case "float32":
		return 32
	case "float64":
		return 64
	default:
		return 0
	}
}
