package <%= package_name %>

import (
	"context"

	tracing "github.com/ricardo-ch/go-tracing"
)

type <%= camel_model_name %>Tracing struct {
	next Service
}

// New<%= model_name %>Tracing ...
func New<%= model_name %>Tracing(i Service) Service {
	return <%= camel_model_name %>Tracing{
		next: i,
	}
}

func (s <%= camel_model_name %>Tracing) Get<%= model_name %>(ctx context.Context, id string) (m *<%= model_name %>, err error) {
	span, ctx := tracing.CreateSpan(ctx, "<%= camel_model_name %>s.service::Get<%= model_name %>", &map[string]interface{}{"<%= camel_model_name %>ID": id})
	defer func() {
		if err != nil {
			tracing.SetSpanError(span, err)
		}
		span.Finish()
	}()

	return s.next.Get<%= model_name %>(ctx, id)
}

func (s <%= camel_model_name %>Tracing) Create<%= model_name %>(ctx context.Context, model <%= model_name %>) (id string, err error) {
	span, ctx := tracing.CreateSpan(ctx, "<%= camel_model_name %>s.service::Create<%= model_name %>", &map[string]interface{}{"<%= camel_model_name %>ID": id})
	defer func() {
		if err != nil {
			tracing.SetSpanError(span, err)
		}
		span.Finish()
	}()

	return s.next.Create<%= model_name %>(ctx, model)
}