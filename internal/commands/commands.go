// Package commands defines CLI commands for the application.
//
// Each command lives in its own sub-package and exposes a Command constructor
// that returns a [*cli.Command]. The commands package aggregates them for the
// root CLI shell in internal/app.go.
package commands

import (
	"github.com/capcom6/go-cli-template/internal/commands/example"
	"github.com/urfave/cli/v3"
)

// Commands returns all available CLI commands.
func Commands() []*cli.Command {
	return []*cli.Command{
		example.Command(),
	}
}
