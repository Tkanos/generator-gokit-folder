package <%= package_name %>

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/tracing/opentracing"
	httptransport "github.com/go-kit/kit/transport/http"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

// ErrInvalidBody thrown when the body of a request can not be parsed
var ErrInvalidBody = errors.New("invalid body")

// MakeHTTPHandler returns all http handler for the <%= camel_model_name %> service
func MakeHTTPHandler(endpoints Endpoints, tracer stdopentracing.Tracer) http.Handler {
	options := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	get<%= model_name %>Handler := kithttp.NewServer(
		endpoints.GetByID,
		decodeGet<%= model_name %>Request,
		encodeResponse,
		append(options, httptransport.ServerBefore(opentracing.FromHTTPRequest(tracer, "calling HTTP GET /{id}", nil)))...,
	)

	create<%= model_name %>Handler := kithttp.NewServer(
		endpoints.Create,
		decodeCreate<%= model_name %>Request,
		encodeCreate<%= model_name %>Response,
		append(options, httptransport.ServerBefore(opentracing.FromHTTPRequest(tracer, "calling POST /", nil)))...,
	)

	r := mux.NewRouter().PathPrefix("/<%= camel_model_name %>s/").Subrouter().StrictSlash(true)

	r.Handle("/{id}", get<%= model_name %>Handler).Methods("GET")
	r.Handle("/", create<%= model_name %>Handler).Methods("POST")

	return r
}

func decodeGet<%= model_name %>Request(ctx context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)

	return Get<%= model_name %>Request{ID: vars["id"]}, nil
}

func decodeCreate<%= model_name %>Request(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req Create<%= model_name %>Request

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrInvalidBody
	}

	return req, nil
}


func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeCreate<%= model_name %>Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	// TODO : refactor return
	ID, ok := response.(string)
	if !ok {
		return errors.New("An error occured while creating <%= model_name %>")
	}
	w.Header().Set("Location", fmt.Sprintf("/<%= camel_model_name %>s/%v", ID))
	w.WriteHeader(http.StatusCreated)
	return nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	switch err {
	case ErrInvalidBody:
		w.WriteHeader(http.StatusBadRequest)
	case ErrNotFound:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		//logger.Error("", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
