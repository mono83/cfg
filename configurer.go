package cfg

// Configurer interface describes configuration source
type Configurer interface {
	// Has returns true if configurer has requested field
	Has(string) bool

	// UnmarshalKey writes configuration value from string key into interface target
	UnmarshalKey(string, interface{}) error

	// KeyFunc return Unmarshaling function for requested key
	KeyFunc(string) func(interface{}) error
}
