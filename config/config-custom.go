package config

import "github.com/mono83/cfg"

// Custom methods registers custom configuration source
// from provided map
func Custom(src map[string]interface{}) {
	def.Custom(src)
}

// Custom methods registers custom configuration source
// from provided map
func (c *Config) Custom(src map[string]interface{}) {
	c.configs = append(c.configs, cfg.Map(src))
	c.clear()
}
