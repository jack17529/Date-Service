package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(HelloService) HelloService

type loggingMiddleware struct {
	logger log.Logger
	next   HelloService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a HelloService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next HelloService) HelloService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Status(ctx context.Context) (s0 string, e1 error) {
	defer func() {
		l.logger.Log("method", "Status", "s0", s0, "e1", e1)
	}()
	return l.next.Status(ctx)
}
func (l loggingMiddleware) Get(ctx context.Context) (s0 string, e1 error) {
	defer func() {
		l.logger.Log("method", "Get", "s0", s0, "e1", e1)
	}()
	return l.next.Get(ctx)
}
func (l loggingMiddleware) Validate(ctx context.Context, date string) (b0 bool, e1 error) {
	defer func() {
		l.logger.Log("method", "Validate", "date", date, "b0", b0, "e1", e1)
	}()
	return l.next.Validate(ctx, date)
}
