//+build unit

package <%= package_name %>

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New<%= model_name %>Tracing_Should_Create_New_Service_Instance(t *testing.T) {
	//Arrange
	fake := new(mockedService)

	//Act
	tracer := New<%= model_name %>Tracing(fake)

	//Assert
	assert.NotNil(t, tracer)
}

func Test_Get<%= model_name %>_Should_Go_Throught_The_Method(t *testing.T) {
	//Arrange
	fake := new(mockedService)
	iD := "123"
	fake.On("Get<%= model_name %>", iD).Return(&<%= model_name %>{}, nil)
	tracer := New<%= model_name %>Tracing(fake)

	//Act
	c, err := tracer.Get<%= model_name %>(context.Background(), iD)

	//Assert
	assert.NotNil(t, c)
	assert.NoError(t, err)
}

func Test_Get<%= model_name %>_With_Error_Should_Go_Throught_The_Method(t *testing.T) {
	//Arrange
	fake := new(mockedService)
	iD := "123"
	errorExpected := errors.New("test")
	fake.On("Get<%= model_name %>", iD).Return(&<%= model_name %>{}, errors.New("test"))
	tracer := New<%= model_name %>Tracing(fake)

	//Act
	c, err := tracer.Get<%= model_name %>(context.Background(), iD)

	//Assert
	assert.NotNil(t, c)
	assert.Equal(t, errorExpected.Error(), err.Error(), "Middleware <%= model_name %>s tracing")
}

func Test_Create<%= model_name %>_Should_Go_Throught_The_Method(t *testing.T) {
	//Arrange
	fake := new(mockedService)
	model := <%= model_name %>{
		<%= model_name %>ID : "123",
	} 
	fake.On("Create<%= model_name %>", model).Return(model.<%= model_name %>ID, nil)
	tracer := New<%= model_name %>Tracing(fake)

	//Act
	c, err := tracer.Create<%= model_name %>(context.Background(), model)

	//Assert
	assert.NotNil(t, c)
	assert.NoError(t, err)
}

func Test_Create<%= model_name %>_With_Error_Should_Go_Throught_The_Method(t *testing.T) {
	//Arrange
	fake := new(mockedService)
	model := <%= model_name %>{
		<%= model_name %>ID : "123",
	} 
	fake.On("Create<%= model_name %>", model).Return(model.<%= model_name %>ID, errors.New("test"))
	errorExpected := errors.New("test")
	tracer := New<%= model_name %>Tracing(fake)

	//Act
	c, err := tracer.Create<%= model_name %>(context.Background(), model)

	//Assert
	assert.NotNil(t, c)
	assert.Equal(t, errorExpected.Error(), err.Error(), "Middleware <%= model_name %>s tracing")
}