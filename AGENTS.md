# go-cli-template — AGENTS.md

## Commands
- `make deps` — go mod download
- `make gen` — go generate ./... (template ships no generators)
- `make fmt` — golangci-lint fmt (goimports, golines 120, swaggo formatter)
- `make lint` — golangci-lint run --timeout=5m (very strict config)
- `make test` — go test -race -shuffle=on -count=1 -covermode=atomic -coverpkg=./... -coverprofile=coverage.out ./...
- `make coverage` — test + go tool cover text + HTML
- `make build` — binary to bin/$(BINARY_NAME) (directory name)
- `make air` — live reload (requires air, sets TZ=UTC DEBUG=1)
- `make release` — goreleaser release --snapshot --clean

## Architecture
- **Entrypoint**: main.go — loads `.env` **before** constructing the logger (so `LOG_*`/`NO_COLOR` apply), builds the logger, wires the signal context, installs the version printer/flag, runs the CLI, maps errors to exit codes, closes the logger. Version injected via ldflags (`appVersion`, `appBuildDate`, `appGitCommit`)
- **CLI**: urfave/cli/v3 — root command in internal/app.go (prints help; global `--debug`/`--version`, no root action); commands registered in internal/commands/commands.go
- **Wiring**: plain constructors, no DI framework — each command's action runs `parseConfig` → `Validate` → constructs the domain service with `New(cfg, log)`
- **Config**: environment-only; `.env` via godotenv at startup; flags read env sources via `cli.EnvVars` (e.g. `DEBUG`, `EXAMPLE__VALUE`)
- **Logging**: go-core-fx/cli-logger — `logger.NewDefault()` in main, `logger.GetLogger(ctx)` inside commands
- **Exit codes**: internal/cli/codes — 0 Success, 1 ParamsError, 2 ClientError, 3 OutputError, 4 InternalError. Usage errors from flag parsing are classified as `ParamsError` by `isUsageError` in main.go
- **No** HTTP server, bot, database, or metrics in the template

## Module Conventions
- Command packages expose `Command() *cli.Command` and are listed in internal/commands/commands.go
- `parseConfig(cmd)` maps flags onto the domain `Config`, wraps validation with `fmt.Errorf("%w: …", Err…)`; the action returns `cli.Exit(err.Error(), codes.ParamsError)` on invalid config
- Domain packages live in internal/<name>/ with config.go (Config + Validate), domain.go, errors.go (sentinel `Err*`), service.go (`New(config, logger.Logger)` plain constructor)
- Failed result writes return `cli.Exit(…, codes.OutputError)`

## Code Generation
- `make gen` → `go generate ./...` — add `//go:generate` directives as needed
- Live reload config: .air.toml

## Linting & CI
- golangci-lint v2.14.0 or later, ~70 linters (exhaustruct_v5, gochecknoglobals, godot, err113, wrapcheck, sloglint, etc.)
- Formatters: goimports, golines max length 120, swaggo
- Format + lint + test + coverage in that order
- CI: go.yml — lint + test on push/PR to master (path-filtered to `**.go`/`go.mod`)
- CI: pr.yml — GoReleaser snapshot, Docker image build, artifact upload on PRs
- CI: release.yml — goreleaser release on `v*` tags
- CI: pr-labels.yml removes review labels on new pushes; close-issues-prs.yml closes stale issues/PRs after 14d inactivity
- E2E CI job is disabled (if: false)

## Testing
- No tests exist yet — add _test.go as needed
- Test flags: -race -shuffle=on -count=1 (disables cache, randomizes order)

## Build & Release
- GoReleaser v2 for linux/windows/darwin, CGO_ENABLED=0
- Go 1.25+
- Docker images pushed to ghcr.io: `pr-<n>` tags on PRs, version tags on release
