//+build unit

package <%= package_name %>

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/ricardo-ch/payment-api/logger"

	"github.com/stretchr/testify/assert"
)

func Test_MakeHTTPHandler(t *testing.T) {
	h := MakeHTTPHandler(Endpoints{}, mocktracer.New())

	assert.NotNil(t, h)
}

func Test_DecodeGet<%= model_name %>Request(t *testing.T) {
	expected := Get<%= model_name %>Request{}
	r, _ := http.NewRequest("GET", "/<%= camel_model_name %>s/1", nil)

	req, err := decodeGet<%= model_name %>Request(context.Background(), r)

	assert.Nil(t, err)
	assert.Equal(t, expected, req)
}

func Test_DecodeCreate<%= model_name %>Request(t *testing.T) {
	expected := Create<%= model_name %>Request{<%= model_name %>{<%= model_name %>ID: "123"}}
	r, _ := http.NewRequest("POST", "/<%= camel_model_name %>s/", bytes.NewBufferString("{\"<%= camel_model_name %>_id\":\"123\"}"))

	req, err := decodeCreate<%= model_name %>Request(context.Background(), r)

	assert.Nil(t, err)
	assert.Equal(t, expected, req)
}

func Test_DecodeCreate<%= model_name %>Request_Should_Returns_ErrInvalidBody_When_Body_Is_Invalid(t *testing.T) {
	expected := ErrInvalidBody
	r, _ := http.NewRequest("POST", "/<%= camel_model_name %>s/", bytes.NewBufferString("invalidjson"))

	_, err := decodeCreate<%= model_name %>Request(context.Background(), r)

	assert.Equal(t, expected, err)
}

func Test_EncodeResponse(t *testing.T) {
	response := struct {
		ID string
	}{"12345"}
	expected := "{\"ID\":\"12345\"}\n"

	w := httptest.NewRecorder()
	err := encodeResponse(context.Background(), w, response)
	assert.Nil(t, err)

	body, err := ioutil.ReadAll(w.Body)

	assert.Nil(t, err)
	assert.Equal(t, expected, string(body))
}

func Test_EncodeResponse_Should_Return_JSON_ContentType(t *testing.T) {
	expected := "application/json; charset=utf-8"

	w := httptest.NewRecorder()
	err := encodeResponse(context.Background(), w, nil)

	assert.Nil(t, err)
	assert.Equal(t, expected, w.Header().Get("Content-Type"))
}

func Test_EncodeCreate<%= model_name %>Response(t *testing.T) {
	ID := "123"
	response := ID
	expectedBody := ""
	expectedLocation := fmt.Sprintf("/<%= camel_model_name %>s/%v", ID)

	w := httptest.NewRecorder()
	err := encodeCreate<%= model_name %>Response(context.Background(), w, response)
	assert.Nil(t, err)

	body, err := ioutil.ReadAll(w.Body)

	assert.Nil(t, err)
	assert.Equal(t, expectedBody, string(body))
	assert.Equal(t, expectedLocation, w.Header().Get("Location"))
	assert.Equal(t, http.StatusCreated, w.Code)
}

func Test_EncodeError_Should_Return_JSON_ContentType(t *testing.T) {
	//Arrange
	logger.New()
	expected := "application/json; charset=utf-8"

	w := httptest.NewRecorder()

	//Act
	encodeError(context.Background(), errors.New("error"), w)

	//Assert
	assert.Equal(t, expected, w.Header().Get("Content-Type"))
}

func Test_EncodeError(t *testing.T) {
	err := errors.New("fake error")
	expected := "{\"error\":\"fake error\"}\n"

	w := httptest.NewRecorder()
	encodeError(context.Background(), err, w)
	body, err := ioutil.ReadAll(w.Body)

	assert.Nil(t, err)
	assert.Equal(t, expected, string(body))
}

func Test_EncodeError_Should_Correctly_Map_Error(t *testing.T) {
	var flagtests = []struct {
		in  error
		out int
	}{
		{errors.New("not handled error"), http.StatusInternalServerError},
		{ErrInvalidBody, http.StatusBadRequest},
		{ErrNotFound, http.StatusNotFound},
	}

	for _, tt := range flagtests {
		w := httptest.NewRecorder()
		encodeError(context.Background(), tt.in, w)

		assert.Equal(t, tt.out, w.Code)
	}
}