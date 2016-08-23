package cfg

// ConfigurersList represents slice of configurers and is configurer
// by itself
type ConfigurersList []Configurer

// Has returns true if configurer has requested field
func (c ConfigurersList) Has(key string) bool {
	for _, e := range c {
		if e.Has(key) {
			return true
		}
	}

	return false
}

// UnmarshallKey writes configuration value from string key into interface target
func (c ConfigurersList) UnmarshallKey(key string, target interface{}) error {
	for _, e := range c {
		if e.Has(key) {
			return e.UnmarshallKey(key, target)
		}
	}

	return ErrKeyMissing{Key: key}
}

// KeyFunc return unmarshalling function for requested key
func (c ConfigurersList) KeyFunc(key string) func(interface{}) error {
	return ExtractUnmarshallFunc(c, key)
}
