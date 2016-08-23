package cfg

import "github.com/mono83/cfg/reflect"

// Map is configuration provider, based on simple hash map
type Map map[string]interface{}

// Has returns true if configurer has requested field
func (m Map) Has(key string) bool {
	_, ok := m[key]
	return ok
}

// UnmarshalKey writes configuration value from string key into interface target
func (m Map) UnmarshalKey(key string, target interface{}) error {
	v, ok := m[key]
	if !ok {
		return ErrKeyMissing{Key: key}
	}

	return reflect.CopyHelper(key, v, target)
}

// KeyFunc return Unmarshaling function for requested key
func (m Map) KeyFunc(key string) func(interface{}) error {
	return ExtractUnmarshalFunc(m, key)
}
