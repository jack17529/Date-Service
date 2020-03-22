package service

import (
	"context"
	"time"
)

// HelloService describes the service.
type HelloService interface {
	Status(ctx context.Context) (string, error)
	Get(ctx context.Context) (string, error)
	Validate(ctx context.Context, date string) (bool, error)
}

type basicHelloService struct{}

// NewBasicHelloService returns a naive, stateless implementation of HelloService.
func NewBasicHelloService() HelloService {
	return &basicHelloService{}
}

func (b *basicHelloService) Status(ctx context.Context) (s0 string, e1 error) {
	return "ok", nil
}
func (b *basicHelloService) Get(ctx context.Context) (s0 string, e1 error) {
	now := time.Now()
	return now.Format("02/01/2006"), nil
}
func (b *basicHelloService) Validate(ctx context.Context, date string) (b0 bool, e1 error) {
	_, err := time.Parse("02/01/2006", date)
	if err != nil {
		return false, err
	}
	return true, nil
}

// New returns a HelloService with all of the expected middleware wired in.
func New(middleware []Middleware) HelloService {
	var svc HelloService = NewBasicHelloService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
