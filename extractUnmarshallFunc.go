package cfg

// ExtractUnmarshalFunc is helper function used to create
// Unmarshaling function for provided key
func ExtractUnmarshalFunc(c Configurer, key string) func(interface{}) error {
	return func(r interface{}) error {
		return c.UnmarshalKey(key, r)
	}
}
