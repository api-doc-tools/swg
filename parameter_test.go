package swg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParameter_ToSwaggerParameters(t *testing.T) {
	assert := assert.New(t)
	Parameters := map[string]Parameter{
		"id": Parameter{
			InPath:     &Value{Type: "string"},
			InHeader:   &Value{Type: "string"},
			InQuery:    &Value{Type: "string"},
			InFormData: &Value{Type: "string"},
		},
		"type": Parameter{
			InPath:     &Value{Type: "string", Enum: "TYPE1 TYPE2 TYPE3"},
			InHeader:   &Value{Type: "string", Enum: "TYPE1 TYPE2 TYPE3"},
			InQuery:    &Value{Type: "string", Enum: "TYPE1 TYPE2 TYPE3"},
			InFormData: &Value{Type: "string", Enum: "TYPE1 TYPE2 TYPE3"},
		},
		"file1": Parameter{
			InFormData: &Value{Type: "file"},
		},
	}
	for name, parameter := range Parameters {
		_, err := parameter.ToSwaggerParameters(name)
		assert.Nil(err)
	}

	invalidParameters := map[string]Parameter{
		"file1": Parameter{
			InPath: &Value{Type: "file"}, // InPath  Value.Type cann't be file
		},
		"file2": Parameter{
			InHeader: &Value{Type: "file"}, // InHeader  Value.Type cann't be file
		},
		"file3": Parameter{
			InQuery: &Value{Type: "file"}, // InQuery  Value.Type cann't be file
		},
		"&**&$#%$#%": Parameter{
			InPath: &Value{Type: "string"}, // invalid name
		},
	}

	for name, parameter := range invalidParameters {
		_, err := parameter.ToSwaggerParameters(name)
		assert.NotNil(err)
	}

}
