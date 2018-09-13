package swg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue_checkWithHTTPIn(t *testing.T) {
	assert := assert.New(t)
	v := Value{Type: "file"}
	err := v.checkWithHTTPIn(InFormData)
	assert.Nil(err)
	err = v.checkWithHTTPIn(InHeader)
	assert.NotNil(err)
}

func TestValue_check(t *testing.T) {
	assert := assert.New(t)
	values := []*Value{
		// normal
		&Value{Type: "file"},
		&Value{Type: "string"},
		&Value{Type: "bool"},
		&Value{Type: "int"},
		&Value{Type: "int32"},
		&Value{Type: "int64"},
		&Value{Type: "uint"},
		&Value{Type: "uint32"},
		&Value{Type: "uint64"},
		&Value{Type: "float32"},
		&Value{Type: "float64"},
		// enum
		&Value{Type: "string", Enum: "TYPE1 TYPE2 TYPE3"},
		&Value{Type: "int", Enum: "-100 0 100"},
		&Value{Type: "int32", Enum: "-100 0 100"},
		&Value{Type: "int64", Enum: "-100 0 100"},
		&Value{Type: "uint", Enum: "0 10 100"},
		&Value{Type: "uint32", Enum: "0 10 100"},
		&Value{Type: "uint64", Enum: "0 10 100"},
		// min
		&Value{Type: "int", Min: "-100"},
		&Value{Type: "int32", Min: "-100"},
		&Value{Type: "int64", Min: "-100"},
		&Value{Type: "uint", Min: "1"},
		&Value{Type: "uint32", Min: "1"},
		&Value{Type: "uint64", Min: "1"},
		&Value{Type: "float32", Min: "-5.0"},
		&Value{Type: "float64", Min: "-5.0"},
		// max
		&Value{Type: "int", Max: "100"},
		&Value{Type: "int32", Max: "100"},
		&Value{Type: "int64", Max: "100"},
		&Value{Type: "uint", Max: "99"},
		&Value{Type: "uint32", Max: "99"},
		&Value{Type: "uint64", Max: "99"},
		&Value{Type: "float32", Max: "5.0"},
		&Value{Type: "float64", Max: "5.0"},
		// min and max
		&Value{Type: "int", Min: "-100", Max: "100"},
		&Value{Type: "int32", Min: "-100", Max: "100"},
		&Value{Type: "int64", Min: "-100", Max: "100"},
		&Value{Type: "uint", Min: "1", Max: "99"},
		&Value{Type: "uint32", Min: "1", Max: "99"},
		&Value{Type: "uint64", Min: "1", Max: "99"},
		&Value{Type: "float32", Min: "1", Max: "5.0"},
		&Value{Type: "float64", Min: "1", Max: "5.0"},
	}
	for _, value := range values {
		err := value.check()
		assert.Nil(err)
	}

	invalidValues := []*Value{
		// err: invalid type
		&Value{Type: "xxxx"},
		// err: cann't set enum
		&Value{Type: "file", Enum: "TYPE1 TYPE2 TYPE3"},
		&Value{Type: "float32", Enum: "0.1 0.2 0.3"},
		&Value{Type: "float64", Enum: "0.1 0.2 0.3"},
		// err: invalid enum
		&Value{Type: "int", Enum: "AAA BBB"},
		&Value{Type: "int32", Enum: "AAA BBB"},
		&Value{Type: "int64", Enum: "AAA BBB"},
		&Value{Type: "uint", Enum: "AAA BBB"},
		&Value{Type: "uint32", Enum: "AAA BBB"},
		&Value{Type: "uint64", Enum: "AAA BBB"},
		// err: cann't set min
		&Value{Type: "file", Min: "123"},
		&Value{Type: "string", Min: "a"},
		&Value{Type: "bool", Min: "false"},
		// err: can't set max
		&Value{Type: "file", Max: "123"},
		&Value{Type: "string", Max: "a"},
		&Value{Type: "bool", Max: "false"},
		// err: invalid min
		&Value{Type: "int", Min: "6.66"},
		&Value{Type: "int32", Min: "6.66"},
		&Value{Type: "int64", Min: "6.66"},
		&Value{Type: "uint", Min: "-1"},
		&Value{Type: "uint32", Min: "-1"},
		&Value{Type: "uint64", Min: "-1"},
		&Value{Type: "float32", Min: "abc"},
		&Value{Type: "float64", Min: "abc"},
		// err: invalid max
		&Value{Type: "int", Max: "6.66"},
		&Value{Type: "int32", Max: "6.66"},
		&Value{Type: "int64", Max: "6.66"},
		&Value{Type: "uint", Max: "-1"},
		&Value{Type: "uint32", Max: "-1"},
		&Value{Type: "uint64", Max: "-1"},
		&Value{Type: "float32", Max: "abc"},
		&Value{Type: "float64", Max: "abc"},
		// // err : out of range
		&Value{Type: "int", Max: "99999999999999999999999999999999999999999999999999999"},
		&Value{Type: "int32", Max: "99999999999999999999999999999999999999999999999999999"},
		&Value{Type: "int64", Max: "99999999999999999999999999999999999999999999999999999"},
		&Value{Type: "uint", Max: "99999999999999999999999999999999999999999999999999999"},
		&Value{Type: "uint32", Max: "99999999999999999999999999999999999999999999999999999"},
		&Value{Type: "uint64", Max: "99999999999999999999999999999999999999999999999999999"},
		&Value{Type: "float32", Max: "99999999999999999999999999999999999999999999999999999"},
		// err: min < max
		&Value{Type: "int", Max: "-100", Min: "100"},
		&Value{Type: "int32", Max: "-100", Min: "100"},
		&Value{Type: "int64", Max: "-100", Min: "100"},
		&Value{Type: "uint", Max: "1", Min: "99"},
		&Value{Type: "uint32", Max: "1", Min: "99"},
		&Value{Type: "uint64", Max: "1", Min: "99"},
		&Value{Type: "float32", Max: "1", Min: "5.0"},
		&Value{Type: "float64", Max: "1", Min: "5.0"},
		// err: Enum and range can't coexist
		&Value{Type: "int", Enum: "-100 0 100", Min: "-100"},
		&Value{Type: "int32", Enum: "-100 0 100", Min: "-100"},
		&Value{Type: "int64", Enum: "-100 0 100", Min: "-100"},
		&Value{Type: "uint", Enum: "10 100", Min: "10"},
		&Value{Type: "uint32", Enum: "10 100", Min: "10"},
		&Value{Type: "uint64", Enum: "10 100", Min: "10"},
		&Value{Type: "int", Enum: "-100 0 100", Max: "100"},
		&Value{Type: "int32", Enum: "-100 0 100", Max: "100"},
		&Value{Type: "int64", Enum: "-100 0 100", Max: "100"},
		&Value{Type: "uint", Enum: "10 100", Max: "100"},
		&Value{Type: "uint32", Enum: "10 100", Max: "100"},
		&Value{Type: "uint64", Enum: "10 100", Max: "100"},
	}
	for _, value := range invalidValues {
		err := value.check()
		assert.NotNil(err)
	}
}

func TestValue_toSwaggerHeader(t *testing.T) {
	assert := assert.New(t)
	values := []*Value{
		// normal
		&Value{Type: "string"},
		&Value{Type: "bool"},
		&Value{Type: "int"},
		&Value{Type: "int32"},
		&Value{Type: "int64"},
		&Value{Type: "uint"},
		&Value{Type: "uint32"},
		&Value{Type: "uint64"},
		&Value{Type: "float32"},
		&Value{Type: "float64"},
		// enum
		&Value{Type: "string", Enum: "TYPE1 TYPE2 TYPE3"},
		&Value{Type: "int", Enum: "-100 0 100"},
		&Value{Type: "int32", Enum: "-100 0 100"},
		&Value{Type: "int64", Enum: "-100 0 100"},
		&Value{Type: "uint", Enum: "0 10 100"},
		&Value{Type: "uint32", Enum: "0 10 100"},
		&Value{Type: "uint64", Enum: "0 10 100"},
		// min
		&Value{Type: "int", Min: "-100"},
		&Value{Type: "int32", Min: "-100"},
		&Value{Type: "int64", Min: "-100"},
		&Value{Type: "uint", Min: "1"},
		&Value{Type: "uint32", Min: "1"},
		&Value{Type: "uint64", Min: "1"},
		&Value{Type: "float32", Min: "-5.0"},
		&Value{Type: "float64", Min: "-5.0"},
		// max
		&Value{Type: "int", Max: "100"},
		&Value{Type: "int32", Max: "100"},
		&Value{Type: "int64", Max: "100"},
		&Value{Type: "uint", Max: "99"},
		&Value{Type: "uint32", Max: "99"},
		&Value{Type: "uint64", Max: "99"},
		&Value{Type: "float32", Max: "5.0"},
		&Value{Type: "float64", Max: "5.0"},
		// min and max
		&Value{Type: "int", Min: "-100", Max: "100"},
		&Value{Type: "int32", Min: "-100", Max: "100"},
		&Value{Type: "int64", Min: "-100", Max: "100"},
		&Value{Type: "uint", Min: "1", Max: "99"},
		&Value{Type: "uint32", Min: "1", Max: "99"},
		&Value{Type: "uint64", Min: "1", Max: "99"},
		&Value{Type: "float32", Min: "1", Max: "5.0"},
		&Value{Type: "float64", Min: "1", Max: "5.0"},
	}

	for _, value := range values {
		_, err := value.toSwaggerHeader()
		assert.Nil(err)
	}

	invalidValues := []*Value{
		&Value{Type: "file"},
	}

	for _, value := range invalidValues {
		_, err := value.toSwaggerHeader()
		assert.NotNil(err)
	}
}
