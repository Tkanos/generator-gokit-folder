package <%= package_name %>

import (
	"context"
	"errors"

)

// ErrNotFound is used when an <%= model_name %> is not found
var ErrNotFound = errors.New("<%= model_name %> not found")

// Service is the Order service interface
type Service interface {
	Get<%= model_name %>(ctx context.Context, id string) (*<%= model_name %>, error)
	Create<%= model_name %>(ctx context.Context, model <%= model_name %>) (string, error)
}

type service struct {
	
}

// NewService return a new instance of order service
func NewService() Service {
	return service{}
}

// Get<%= model_name %> returns an <%= model_name %> regarding the id passed in parameter
func (s service) Get<%= model_name %>(ctx context.Context, id string) (m *<%= model_name %>, err error) {

	m = &<%= model_name %>{ <%= model_name %>ID : id}
	return 
}

// Create<%= model_name %> creates an <%= model_name %>
func (s service) Create<%= model_name %>(ctx context.Context, model <%= model_name %>) (id string, err error) {

	id = model.<%= model_name %>ID
	return
}
