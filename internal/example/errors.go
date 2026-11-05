package example

import "errors"

// Sentinel errors for the example module. Callers wrap them with
// [fmt.Errorf]("%w: detail", Err...) so application logic can branch
// with [errors.Is].
var (
	// ErrConfigInvalid indicates the module configuration cannot be used.
	ErrConfigInvalid = errors.New("invalid config")

	// ErrExample is a predefined error for the example module.
	ErrExample = errors.New("example error")
)
