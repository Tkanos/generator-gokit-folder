//+build unit

package <%= package_name %>

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MakeGetByIDEndpoint(t *testing.T) {
	fakeService := new(mockedService)
	fakeService.On("Get<%= model_name %>", "1").Return(&(<%= model_name %>{}), nil)

	e := MakeGetByIDEndpoint(fakeService)
	u, err := e(nil, Get<%= model_name %>Request{ID: "1"})

	assert.NotNil(t, e)
	assert.Nil(t, err)
	assert.NotNil(t, u)
}

func Test_MakeCreateEndpoint(t *testing.T) {
	fakeService := new(mockedService)
	fakeService.On("Create<%= model_name %>", <%= model_name %>{}).Return(*new(string), nil)

	e := MakeCreateEndpoint(fakeService)
	_, err := e(nil, Create<%= model_name %>Request{})

	assert.NotNil(t, e)
	assert.Nil(t, err)
}