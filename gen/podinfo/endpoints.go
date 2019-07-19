// Code generated by goa v3.0.3, DO NOT EDIT.
//
// podinfo endpoints
//
// Command:
// $ goa gen github.com/centric-lt/k8s-101/design

package podinfo

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "podinfo" service endpoints.
type Endpoints struct {
	Get goa.Endpoint
}

// NewEndpoints wraps the methods of the "podinfo" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Get: NewGetEndpoint(s),
	}
}

// Use applies the given middleware to all the "podinfo" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Get = m(e.Get)
}

// NewGetEndpoint returns an endpoint function that calls the method "get" of
// service "podinfo".
func NewGetEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		res, err := s.Get(ctx)
		if err != nil {
			return nil, err
		}
		vres := NewViewedPodinforesult(res, "default")
		return vres, nil
	}
}
