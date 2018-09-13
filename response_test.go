package swg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse_ToSwaggerResponse(t *testing.T) {
	assert := assert.New(t)
	type LoginInfo struct {
		ID string `json:"id" xml:"id"`
	}
	resp := &Response{
		Description: "return a book",
		Model:       &LoginInfo{},
		Headers: map[string]Value{
			"X-Rate-Limit":    Value{Type: "int32", Desc: "calls per hour allowed by the user"},
			"X-Expires-After": Value{Type: "string", Desc: "date in UTC when token expires"},
		},
	}
	_, err := resp.ToSwaggerResponse()
	assert.Nil(err)

	// err: invalid Value.Type in Response.Headers
	invalidResp := &Response{
		Description: "return a book",
		Model:       &LoginInfo{},
		Headers: map[string]Value{
			"X-Rate-Limit":    Value{Type: "int32", Desc: "calls per hour allowed by the user"},
			"X-Expires-After": Value{Type: "file", Desc: "date in UTC when token expires"}, // Response.Headers Value.Type cann't be file
		},
	}
	_, err = invalidResp.ToSwaggerResponse()
	assert.NotNil(err)

	// err: invalid model
	invalidResp = &Response{
		Description: "return a book",
		Model:       map[string]string{}, // model must be a struct or []struct
	}
	_, err = invalidResp.ToSwaggerResponse()
	assert.NotNil(err)
}
