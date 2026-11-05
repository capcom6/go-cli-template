# Scaffolding a New Go CLI from This Template

## Initial Setup

Two paths:

**GitHub template** — click "Use this template" on the repo page, clone your new repo, `cd` in.

**Manual** — clone and re-init:
```
git clone git@github.com:capcom6/go-cli-template.git <your-project>
cd <your-project>
rm -rf .git
git init
git add .
git commit -m "init from go-cli-template"
```

## Replace Placeholder Values

Search-and-replace table:

| Find | Replace with |
|---|---|
| `github.com/capcom6/go-cli-template` | `your/module/path` (in `go.mod`, all `.go` imports) |
| `go-cli-template` | your binary name (root `Command.Name` in `internal/app.go`; `Makefile` derives `BINARY_NAME` from the directory name) |

**`internal/app.go`** — update the root command:
- `Name`, `Usage`, `Description`
- `Authors`

**`.goreleaser.yaml`** — update:
- `DOCKER_REGISTRY` env / labels `source`, `name`

Version globals in `main.go` stay; they get real values via goreleaser ldflags
(`appVersion`, `appBuildDate`, `appGitCommit`).

## Remove (or Adapt) the Example Command

1. `internal/commands/example/` — delete the directory (command + flag parsing)
2. `internal/example/` — delete the directory (business domain)
3. `internal/commands/commands.go` — remove `example.Command()` from the list and drop the `example` import

## Adding a New Command

Each command is a plain constructor — no DI framework:

```
internal/commands/<name>/
├── config.go     # maps flags onto the domain Config + validation
└── <name>.go     # Command() *cli.Command + action
```

**`<name>.go`** template:

```go
// Command returns a one-shot <name> command.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "<name>",
		Usage: "…",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "<flag>",
				Usage:   "…",
				Sources: cli.EnvVars("<NAME>__<FLAG>"),
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

	svc := <name>.New(cfg, logger.GetLogger(ctx))
	// … do the work …

	return nil
}
```

Register it in `internal/commands/commands.go`:

```go
func Commands() []*cli.Command {
	return []*cli.Command{
		example.Command(),
		<name>.Command(),
	}
}
```

### Business domain package

For heavier domains add `internal/<name>/`:

```
internal/<name>/
├── config.go     # Config struct with Validate()
├── domain.go     # domain entities
├── errors.go     # sentinel errors (Err*)
└── service.go    # business logic, New(config, log)
```

- Constructor: `func New(cfg Config, log logger.Logger) *Service` — plain
  parameters, constructed by the command layer.
- Wrap sentinel errors with `fmt.Errorf("%w: detail", Err…)` and branch with
  `errors.Is`.
- See `internal/example/` for a complete reference.

### Exit codes

`internal/cli/codes` defines the process exit codes:

| Code | Constant | Meaning |
|---|---|---|
| 0 | `Success` | command completed |
| 1 | `ParamsError` | invalid flags, arguments, or configuration |
| 2 | `ClientError` | failure caused by an external system |
| 3 | `OutputError` | result could not be produced or written |
| 4 | `InternalError` | unexpected internal failure |

- Invalid configuration: `cli.Exit(err.Error(), codes.ParamsError)` — urfave/cli
  exits with the code directly.
- Usage errors from flag parsing are mapped to `ParamsError` in `main.go`
  (`isUsageError`); everything else falls back to `InternalError`.
- Failed writes: `cli.Exit(…, codes.OutputError)`.

## Local Development

```
make deps
make air       # live reload, requires `air` CLI: go install github.com/air-verse/air@latest
```

CLI usage:
```
make build && ./bin/<your-binary-name> example    # one-shot demo (name = directory name)
./bin/<your-binary-name> example --value foo      # flag (env: EXAMPLE__VALUE=foo)
./bin/<your-binary-name> --help                   # list commands
./bin/<your-binary-name> --version                # print version
```

Configuration is environment-only:
- `.env` is loaded at startup **before** the logger is constructed, so
  `LOG_LEVEL`, `LOG_FORMAT`, `LOG_OUTPUT`, and `NO_COLOR` from `.env` apply
  (see `.env.example`).
- Every flag can be fed from an environment variable via `cli.EnvVars`.

## Common Gotchas

- **`exhaustruct`** — every struct literal must initialize all fields; use named
  field syntax, never zero-value shorthand.
- **`gochecknoglobals`** — bans package-level `var`/`const` except build
  metadata (annotate with `//nolint:gochecknoglobals`).
- **`sloglint` `no-global: all`** — global loggers are forbidden; use the
  injected `logger.Logger` from `go-core-fx/cli-logger`.
- **dotenv order** — keep `godotenv.Load()` before `logger.NewDefault()` in
  `main.go`, otherwise `.env` logger settings are silently ignored.
- **`-count=1`** in the test command disables Go test cache. Tests always run
  fresh with randomized order (`-shuffle=on`).

## Build & Deploy

```
make build              # binary to bin/<dirname>
make release            # goreleaser snapshot (local)
git tag v0.1.0 && git push origin v0.1.0   # triggers CI release + Docker push to ghcr.io
```

GoReleaser builds linux/windows/darwin with `CGO_ENABLED=0`. Docker images are
pushed to `ghcr.io` as `pr-<n>` tags on pull requests and version tags on
release.
