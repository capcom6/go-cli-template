package example

import (
	"context"
	"fmt"

	"github.com/capcom6/go-cli-template/internal/cli/codes"
	"github.com/capcom6/go-cli-template/internal/example"
	logger "github.com/go-core-fx/cli-logger"
	"github.com/urfave/cli/v3"
)

// Command returns a one-shot command that demonstrates flag-driven
// configuration and a service wired with plain constructors.
//
// This pattern is useful for CLI tools, migrations, data sync, admin tasks, etc.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "example",
		Usage: "Run an example one-shot task",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "value",
				Usage:   "value to echo back",
				Value:   "example",
				Sources: cli.EnvVars("EXAMPLE__VALUE"),
			},
		},
		Action: run,
	}
}

func run(ctx context.Context, cmd *cli.Command) error {
	cfg, err := parseConfig(cmd)
	if err != nil {
		return cli.Exit(err.Error(), codes.ParamsError)
	}

	log := logger.GetLogger(ctx)
	svc := example.New(cfg, log)

	if _, writeErr := fmt.Fprintf(cmd.Root().Writer, "example value: %s\n", svc.Example(ctx)); writeErr != nil {
		return cli.Exit(fmt.Sprintf("failed to write result: %v", writeErr), codes.OutputError)
	}

	return nil
}
