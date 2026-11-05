package example

import (
	"fmt"

	"github.com/capcom6/go-cli-template/internal/example"
	"github.com/urfave/cli/v3"
)

// parseConfig maps the example command flags onto the example module config
// and validates the result.
func parseConfig(cmd *cli.Command) (example.Config, error) {
	cfg := example.Config{
		Example: cmd.String("value"),
	}

	if err := cfg.Validate(); err != nil {
		return example.Config{}, fmt.Errorf("failed to parse example configuration: %w", err)
	}

	return cfg, nil
}
