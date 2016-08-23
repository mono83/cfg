package file

import (
	"errors"
	"fmt"
	"github.com/mono83/cfg"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

// New creates file-based configuration source
func New(fileNames []string, read func(string) ([]byte, error), build func([]byte) (cfg.Configurer, error)) *Configurer {
	if read == nil {
		read = ioutil.ReadFile
	}

	return &Configurer{
		fileNames:   fileNames,
		fileReader:  read,
		configBuild: build,
	}
}

// Configurer is container for file-based configuration
type Configurer struct {
	fileNames   []string
	fileReader  func(string) ([]byte, error)
	configBuild func([]byte) (cfg.Configurer, error)

	currentConfig cfg.Configurer
	m             sync.Mutex
}

// Reload reloads file container
func (c *Configurer) Reload() error {
	c.m.Lock()
	defer c.m.Unlock()

	// Searching for file
	found := ""
	for _, name := range c.fileNames {
		_, err := os.Stat(name)

		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return err
		}

		found = name
		break
	}

	if found == "" {
		return fmt.Errorf("Unable to locate configuration. Searched in %s", strings.Join(c.fileNames, ","))
	}

	bts, err := c.fileReader(found)
	if err != nil {
		return err
	}

	c.currentConfig, err = c.configBuild(bts)
	return err
}

// Validate validates container
func (c *Configurer) Validate() error {
	if c.currentConfig != nil {
		return nil
	}

	return c.Reload()
}

// Has returns true if configurer has requested field
func (c *Configurer) Has(key string) bool {
	if c.currentConfig == nil {
		return false
	}

	return c.currentConfig.Has(key)
}

// UnmarshalKey writes configuration value from string key into interface target
func (c *Configurer) UnmarshalKey(key string, target interface{}) error {
	if c.currentConfig == nil {
		return errors.New("File source config not ready")
	}

	return c.currentConfig.UnmarshalKey(key, target)
}

// KeyFunc return Unmarshaling function for requested key
func (c *Configurer) KeyFunc(key string) func(interface{}) error {
	return cfg.ExtractUnmarshalFunc(c, key)
}
