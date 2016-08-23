package cfg

// ConfigurerWithAliases is special configurer adapter, that works with aliases
type ConfigurerWithAliases struct {
	aliases map[string]string
	real    Configurer
}

// NewConfigurationWithAliases creates new configurere with empty aliases list
func NewConfigurationWithAliases(real Configurer) *ConfigurerWithAliases {
	return &ConfigurerWithAliases{
		aliases: map[string]string{},
		real:    real,
	}
}

// Alias registers new alias
func (c *ConfigurerWithAliases) Alias(virtual, real string) {
	c.aliases[virtual] = real
}

// Has returns true if configurer has requested field
func (c ConfigurerWithAliases) Has(key string) bool {
	if c.real.Has(key) {
		return true
	}

	v, ok := c.aliases[key]
	return ok && c.real.Has(v)
}

// UnmarshallKey writes configuration value from string key into interface target
func (c ConfigurerWithAliases) UnmarshallKey(key string, target interface{}) error {
	if !c.real.Has(key) {
		v, ok := c.aliases[key]
		if !ok || !c.real.Has(v) {
			return ErrKeyMissing{Key: key}
		}

		key = v
	}

	return c.real.UnmarshallKey(key, target)
}

// KeyFunc return unmarshalling function for requested key
func (c ConfigurerWithAliases) KeyFunc(key string) func(interface{}) error {
	return ExtractUnmarshallFunc(c, key)
}
