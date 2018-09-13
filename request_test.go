package swg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest_toSwaggerParameter(t *testing.T) {
	assert := assert.New(t)
	type book struct {
		ID string
	}
	req := &Request{
		Description: "book",
		Model:       &book{},
	}
	_, err := req.toSwaggerParameter()
	assert.Nil(err)

	invalidReq := &Request{
		Description: "book",
		Model:       nil,
	}
	_, err = invalidReq.toSwaggerParameter()
	assert.NotNil(err)
}
