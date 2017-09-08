//+build unit

package <%= package_name %>

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type mockedService struct {
	mock.Mock
}

func (m *mockedService) Get<%= model_name %>(ctx context.Context, id string) (*<%= model_name %>, error) {
	args := m.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*<%= model_name %>), args.Error(1)
}

func (m *mockedService) Create<%= model_name %>(ctx context.Context, model <%= model_name %>) (string, error) {
	args := m.Called(model)
	return args.Get(0).(string), args.Error(1)
}

