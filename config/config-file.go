package config

import (
	"fmt"
	"path/filepath"
	"strings"
)

// AutoFind registers configuration source read from required file
// File will searched in current folder, then in home folder and then in /etc/
// File type will be determined from its extension
func AutoFind(name, subfolderName string) error {
	return def.AutoFind(name, subfolderName)
}

// AutoFind registers configuration source read from required file
// File will searched in current folder, then in home folder and then in /etc/
// File type will be determined from its extension
func (c *Config) AutoFind(name, subfolderName string) error {
	ext := filepath.Ext(name)
	ext = strings.ToLower(ext)
	switch ext {
	case ".yaml", ".yml":
		c.YAMLAutoFind(name, subfolderName)
	case ".json":
		c.JSONAutoFind(name, subfolderName)
	case "":
		return fmt.Errorf("File %s has no extension and it's type cannot be detemined", name)
	default:
		return fmt.Errorf("Extension %s for file %s not supported", ext, name)
	}

	return nil
}
