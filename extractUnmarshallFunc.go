package cfg

// ExtractUnmarshallFunc is helper function used to create
// unmarshalling function for provided key
func ExtractUnmarshallFunc(c Configurer, key string) func(interface{}) error {
	return func(r interface{}) error {
		return c.UnmarshallKey(key, r)
	}
}
