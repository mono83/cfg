package cfg

// List represents slice of configurers and is configurer
// by itself
// Configurers in list are processed backwards, so last one
// has maximum priority
type List []Configurer

// Has returns true if configurer has requested field
func (c List) Has(key string) bool {
	for _, e := range c {
		if e.Has(key) {
			return true
		}
	}

	return false
}

// UnmarshalKey writes configuration value from string key into interface target
func (c List) UnmarshalKey(key string, target interface{}) error {
	for i := len(c) - 1; i >= 0; i-- {
		e := c[i]
		if e.Has(key) {
			return e.UnmarshalKey(key, target)
		}
	}

	return ErrKeyMissing{Key: key}
}

// KeyFunc return Unmarshaling function for requested key
func (c List) KeyFunc(key string) func(interface{}) error {
	return ExtractUnmarshalFunc(c, key)
}
