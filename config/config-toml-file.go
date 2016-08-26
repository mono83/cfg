package config

import (
	"github.com/mono83/cfg/file"
	"github.com/mono83/cfg/toml"
)

// TOMLFile registers configuration source read from TOML file
func TOMLFile(names ...string) {
	def.TOMLFile(names...)
}

// TOMLFile registers configuration source read from TOML file
func (c *Config) TOMLFile(names ...string) {
	c.AddLast(file.New(names, nil, toml.NewBytesSource))
}

// TOMLAutoFind registers configuration source read from required file
// File will searched in current folder, then in home folder and then in /etc/
func TOMLAutoFind(filename, subfolderName string) {
	def.TOMLAutoFind(filename, subfolderName)
}

// TOMLAutoFind registers configuration source read from required file
// File will searched in current folder, then in home folder and then in /etc/
func (c *Config) TOMLAutoFind(filename, subfolderName string) {
	if subfolderName == "" {
		c.TOMLFile(file.CommonFolders(filename)...)
	} else {
		c.TOMLFile(file.CommonFoldersWithSubfolder(filename, subfolderName)...)
	}
}
