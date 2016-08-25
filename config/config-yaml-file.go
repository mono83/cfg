package config

import (
	"github.com/mono83/cfg/file"
	"github.com/mono83/cfg/yaml"
)

// YAMLFile registers configuration source read from YAML file
func YAMLFile(names ...string) {
	def.YAMLFile(names...)
}

// YAMLFile registers configuration source read from YAML file
func (c *Config) YAMLFile(names ...string) {
	c.AddLast(file.New(names, nil, yaml.NewBytesSource))
}

// YAMLAutoFind registers configuration source read from required file
// File will searched in current folder, then in home folder and then in /etc/
func YAMLAutoFind(filename, subfolderName string) {
	def.YAMLAutoFind(filename, subfolderName)
}

// YAMLAutoFind registers configuration source read from required file
// File will searched in current folder, then in home folder and then in /etc/
func (c *Config) YAMLAutoFind(filename, subfolderName string) {
	if subfolderName == "" {
		c.YAMLFile(file.CommonFolders(filename)...)
	} else {
		c.YAMLFile(file.CommonFoldersWithSubfolder(filename, subfolderName)...)
	}
}
