package config

import (
	"github.com/mono83/cfg"
	"github.com/mono83/cfg/file"
	"io/ioutil"
	"sync"
)

// def contains default static instance of Config struct
var def = New()

// New creates new Config instance
func New() *Config {
	return &Config{
		aliases: map[string]string{},
	}
}

// Config is helper configuration structure
type Config struct {
	configs        []cfg.Configurer
	aliases        map[string]string
	withValidation bool

	fileReader func(string) ([]byte, error)
	staging    cfg.Configurer

	cache cfg.Configurer
	m     sync.Mutex
}

// AddFirst registers configuration source and adds it to the
// beginning of configuration sources list, setting min priority
func AddFirst(cc cfg.Configurer) {
	def.AddFirst(cc)
}

// AddFirst registers configuration source and adds it to the
// beginning of configuration sources list, setting min priority
func (c *Config) AddFirst(cc cfg.Configurer) {
	if cc != nil {
		if len(c.configs) == 0 {
			c.configs = []cfg.Configurer{cc}
		} else {
			c.configs = append([]cfg.Configurer{cc}, c.configs...)
		}
	}

	c.clear()
}

// AddLast registers configuration source and adds it to the
// end of configuration sources list, setting max priority
func AddLast(cc cfg.Configurer) {
	def.AddLast(cc)
}

// AddLast registers configuration source and adds it to the
// end of configuration sources list, setting max priority
func (c *Config) AddLast(cc cfg.Configurer) {
	if cc != nil {
		c.configs = append(c.configs, cc)
	}
	c.clear()
}

// Alias method registers configuration key alias
func Alias(virtual, real string) {
	def.Alias(virtual, real)
}

// Alias method registers configuration key alias
func (c *Config) Alias(virtual, real string) {
	c.aliases[virtual] = real
	c.clear()
}

// EnableValidation enables config values validation
func EnableValidation() {
	def.EnableValidation()
}

// EnableValidation enables config values validation
func (c *Config) EnableValidation() {
	c.withValidation = true
	c.clear()
}

// EnablePlaceholdersInFile enables placeholders parsing
func EnablePlaceholdersInFile() {
	def.EnablePlaceholdersInFile()
}

// EnablePlaceholdersInFile enables placeholders parsing
func (c *Config) EnablePlaceholdersInFile() {
	c.fileReader = func(name string) ([]byte, error) {
		return file.PlaceholdersReader(
			ioutil.ReadFile,
			c.getStagingConfigurer(),
		)(name)
	}
}

func (c *Config) getStagingConfigurer() cfg.Configurer {
	return c.staging
}

// clear drops config cache
func (c *Config) clear() {
	c.m.Lock()
	defer c.m.Unlock()

	c.cache = nil
}

type reloader interface {
	Reload() error
}

// real builds and returns configurer object
func (c *Config) real() cfg.Configurer {
	c.m.Lock()
	defer c.m.Unlock()

	if c.cache == nil {
		inProgress := []cfg.Configurer{}
		for _, cc := range c.configs {
			c.staging = cfg.List(inProgress)
			if rcc, ok := cc.(reloader); ok {
				err := rcc.Reload()
				if err != nil {
					c.cache = cfg.ErrConfigurer(err)
					return c.cache
				}
			}
			inProgress = append(inProgress, cc)
		}

		c.cache = cfg.List(c.configs)
		if len(c.aliases) > 0 {
			ac := cfg.NewConfigurationWithAliases(c.cache)
			for k, v := range c.aliases {
				ac.Alias(k, v)
			}

			c.cache = ac
		}

		if c.withValidation {
			c.cache = cfg.NewValidableConfig(c.cache)
		}
	}

	return c.cache
}

// Has returns true if configurer has requested field
func Has(key string) bool {
	return def.Has(key)
}

// Has returns true if configurer has requested field
func (c *Config) Has(key string) bool {
	return c.real().Has(key)
}

// UnmarshalKey writes configuration value from string key into interface target
func UnmarshalKey(key string, target interface{}) error {
	return def.UnmarshalKey(key, target)
}

// UnmarshalKey writes configuration value from string key into interface target
func (c *Config) UnmarshalKey(key string, target interface{}) error {
	return c.real().UnmarshalKey(key, target)
}

// KeyFunc return Unmarshaling function for requested key
func KeyFunc(key string) func(interface{}) error {
	return def.KeyFunc(key)
}

// KeyFunc return Unmarshaling function for requested key
func (c *Config) KeyFunc(key string) func(interface{}) error {
	return c.real().KeyFunc(key)
}
