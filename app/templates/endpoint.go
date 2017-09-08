package <%= package_name %>

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints represent the order service endpoints
type Endpoints struct {
	GetByID endpoint.Endpoint
	Create  endpoint.Endpoint
}

// MakeGetByIDEndpoint returns an endpoint used for getting one <%= camel_model_name %>
func MakeGetByIDEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Get<%= model_name %>Request)

		return s.Get<%= model_name %>(ctx, req.ID)
	}
}

// MakeCreateEndpoint returns an endpoint used for creating a <%= camel_model_name %>
func MakeCreateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Create<%= model_name %>Request)

		return s.Create<%= model_name %>(ctx, req.<%= model_name %>)
	}
}


// Get<%= model_name %>Request represents the request parameters used for getting one <%= model_name %>
type Get<%= model_name %>Request struct {
	ID string `json:"id"`
}

// Create<%= model_name %>Request represents the request parameters used for creating <%= model_name %>
type Create<%= model_name %>Request struct {
	<%= model_name %>
}