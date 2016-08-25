package config

import "github.com/mono83/cfg"

// CustomMap methods registers custom configuration source
// from provided map
func CustomMap(src map[string]interface{}) {
	def.CustomMap(src)
}

// CustomMap methods registers custom configuration source
// from provided map
func (c *Config) CustomMap(src map[string]interface{}) {
	c.AddLast(cfg.Map(src))
}
