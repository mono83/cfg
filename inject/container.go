package inject

import "github.com/mono83/cfg"

// Container represents dependency injection container
type Container interface {
	cfg.Configurer

	// OnAfterInit register function, that will be invoked after service
	// initialization and before it will be cached and returned
	OnAfterInit(func(name string, target interface{}) error)

	// HasService returns true if service with provided name is declared
	HasService(name string) bool
	// GetService searches for service by name and writes it into target
	GetService(name string, target interface{}) error
	// Define defines new service with build function for it
	Define(name string, build func(Container) (interface{}, error))

	// MustGetService is alias for GetService, that panics on error
	MustGetService(name string, target interface{})
}
