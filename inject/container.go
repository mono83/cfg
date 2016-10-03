package inject

import "github.com/mono83/cfg"

// Container represents dependency injection container
type Container interface {
	cfg.Configurer

	// HasService returns true if service with provided name is declared
	HasService(name string) bool
	// GetService searches for service by name and writes it into target
	GetService(name string, target interface{}) error
	// GetServices handles multiple service retrieval request
	GetServices(defs ...Definition) error
	// Define defines new service with build function for it
	Define(name string, build func(Container) (interface{}, error))
}
