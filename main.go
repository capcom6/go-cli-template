package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"

	"github.com/capcom6/go-cli-template/internal"
	"github.com/capcom6/go-cli-template/internal/cli/codes"
	logger "github.com/go-core-fx/cli-logger"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v3"
)

//nolint:gochecknoglobals // build metadata
var (
	appVersion   = "dev"
	appBuildDate = "unknown"
	appGitCommit = "unknown"
	appGoVersion = runtime.Version()
)

func main() {
	// Load environment variables before the logger reads its configuration.
	// The parser error is deliberately not printed: godotenv messages can
	// include the rest of the .env file from the failing line onward.
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, "failed to load .env file")
	}

	log := logger.NewDefault()

	rootCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	ctx := logger.WithLogger(rootCtx, log)

	//nolint:reassign // urfave/cli specific
	cli.VersionPrinter = func(cmd *cli.Command) {
		fmt.Fprintf(cmd.Root().Writer, "Version:    %s\n", appVersion)
		fmt.Fprintf(cmd.Root().Writer, "Build Date: %s\n", appBuildDate)
		fmt.Fprintf(cmd.Root().Writer, "Git Commit: %s\n", appGitCommit)
		fmt.Fprintf(cmd.Root().Writer, "Go Version:  %s\n", appGoVersion)
	}

	//nolint:reassign // urfave/cli specific
	cli.VersionFlag = &cli.BoolFlag{
		Name:        "version",
		Usage:       "print the version",
		HideDefault: true,
		Local:       true,
	}

	exitCode := codes.Success
	if err := internal.Run(ctx, log, appVersion); err != nil {
		log.Error(ctx, "Application failed", err)
		exitCode = codes.InternalError
		if isUsageError(err) {
			exitCode = codes.ParamsError
		}
	}

	if err := log.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to close logger: %v\n", err)
	}

	stop()
	os.Exit(exitCode)
}

// isUsageError reports whether err comes from urfave/cli flag or argument
// parsing, i.e. a correctable input error rather than an internal failure.
//
// urfave/cli v3.10.1 returns plain untyped errors from the parse path, so
// classification is by the message shapes it emits.
func isUsageError(err error) bool {
	msg := err.Error()

	switch {
	case strings.Contains(msg, "flag provided but not defined: "),
		strings.Contains(msg, "flag needs an argument: "):
		return true
	case strings.Contains(msg, "invalid value ") &&
		(strings.Contains(msg, " for flag -") || strings.Contains(msg, " for argument ")):
		return true
	case strings.Contains(msg, "could not parse ") && strings.Contains(msg, " for flag "):
		return true
	default:
		return false
	}
}
