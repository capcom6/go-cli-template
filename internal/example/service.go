package example

import (
	"context"

	logger "github.com/go-core-fx/cli-logger"
)

// Service implements the business logic for the example module.
//
// Dependencies are plain constructor parameters: the command layer
// (internal/commands/example) constructs the service directly, with no
// DI framework involved.
type Service struct {
	// config holds the configuration for the service.
	config Config

	// logger is used for structured logging.
	logger logger.Logger
}

// New creates and initializes a new Service instance.
//
// A constructor may return an error to fail fast on unusable configuration;
// this example validates in the command layer instead, before construction.
func New(config Config, log logger.Logger) *Service {
	return &Service{
		config: config,
		logger: log,
	}
}

// Example returns the example value from the configuration.
func (s *Service) Example(ctx context.Context) string {
	s.logger.Debug(ctx, "building example response")
	s.logger.Info(ctx, "returning example value", logger.Fields{"value": s.config.Example})

	return s.config.Example
}
