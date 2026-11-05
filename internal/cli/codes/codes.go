// Package codes defines the process exit codes returned by the CLI.
package codes

// Exit codes returned by the application.
const (
	// Success indicates the command completed without errors.
	Success = iota
	// ParamsError indicates invalid flags, arguments, or configuration.
	ParamsError
	// ClientError indicates a failure caused by an external system.
	ClientError
	// OutputError indicates results could not be produced or written.
	OutputError
	// InternalError indicates an unexpected internal failure.
	InternalError
)
