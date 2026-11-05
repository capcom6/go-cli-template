package example

import "fmt"

// Config holds the configuration for the example module.
//
// Module configs are plain structs: the command layer maps CLI flags onto
// them (internal/commands/example/config.go) and Validate is called there,
// so invalid configuration fails before the service is constructed.
//
// Example:
//
//	cfg := example.Config{
//	    Example: "demo-value",
//	}
//
//	if err := cfg.Validate(); err != nil {
//	    return err
//	}
type Config struct {
	// Example is the value the module echoes back. In a real module this
	// would be feature flags, endpoints, timeouts, or similar settings.
	Example string
}

// Validate reports whether the configuration can be used.
func (c Config) Validate() error {
	if c.Example == "" {
		return fmt.Errorf("%w: example must not be empty", ErrConfigInvalid)
	}

	return nil
}
