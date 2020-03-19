package endpoint

import (
	"context"

	service "github.com/faith/goKitAdvanced4/hello/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// StatusRequest collects the request parameters for the Status method.
type StatusRequest struct{}

// StatusResponse collects the response parameters for the Status method.
type StatusResponse struct {
	S0 string `json:"status"`
	E1 error  `json:"e1,omitempty"`
}

// MakeStatusEndpoint returns an endpoint that invokes Status on the service.
func MakeStatusEndpoint(s service.HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		s0, e1 := s.Status(ctx)
		return StatusResponse{
			E1: e1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r StatusResponse) Failed() error {
	return r.E1
}

// GetRequest collects the request parameters for the Get method.
type GetRequest struct{}

// GetResponse collects the response parameters for the Get method.
type GetResponse struct {
	S0 string `json:"date"`
	E1 error  `json:"e1,omitempty"`
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		s0, e1 := s.Get(ctx)
		return GetResponse{
			E1: e1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r GetResponse) Failed() error {
	return r.E1
}

// ValidateRequest collects the request parameters for the Validate method.
type ValidateRequest struct {
	Date string `json:"date"`
}

// ValidateResponse collects the response parameters for the Validate method.
type ValidateResponse struct {
	B0 bool  `json:"valid"`
	E1 error `json:"e1,omitempty"`
}

// MakeValidateEndpoint returns an endpoint that invokes Validate on the service.
func MakeValidateEndpoint(s service.HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ValidateRequest)
		b0, e1 := s.Validate(ctx, req.Date)
		return ValidateResponse{
			B0: b0,
			E1: e1,
		}, nil
	}
}

// Failed implements Failer.
func (r ValidateResponse) Failed() error {
	return r.E1
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Status implements Service. Primarily useful in a client.
func (e Endpoints) Status(ctx context.Context) (s0 string, e1 error) {
	request := StatusRequest{}
	response, err := e.StatusEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(StatusResponse).S0, response.(StatusResponse).E1
}

// Get implements Service. Primarily useful in a client.
func (e Endpoints) Get(ctx context.Context) (s0 string, e1 error) {
	request := GetRequest{}
	response, err := e.GetEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetResponse).S0, response.(GetResponse).E1
}

// Validate implements Service. Primarily useful in a client.
func (e Endpoints) Validate(ctx context.Context, date string) (b0 bool, e1 error) {
	request := ValidateRequest{Date: date}
	response, err := e.ValidateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ValidateResponse).B0, response.(ValidateResponse).E1
}
