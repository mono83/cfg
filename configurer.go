package cfg

// Configurer interface describes configuration source
type Configurer interface {
	// Has returns true if configurer has requested field
	Has(string) bool

	// UnmarshallKey writes configuration value from string key into interface target
	UnmarshallKey(string, interface{}) error

	// KeyFunc return unmarshalling function for requested key
	KeyFunc(string) func(interface{}) error
}
