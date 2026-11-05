// Package example demonstrates the module layout of a CLI-oriented business
// domain wired with plain constructors (no DI framework).
//
// Files:
//
//   - config.go: Config struct with Validate()
//   - domain.go: domain entities
//   - errors.go: sentinel errors (Err*)
//   - service.go: business logic, constructed via New()
//
// The command layer (internal/commands/example) maps CLI flags onto Config,
// validates it, and constructs the service. Persistence (repository.go, when
// needed) follows the same plain-constructor style.
package example
