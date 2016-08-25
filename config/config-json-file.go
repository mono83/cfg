package config

import (
	"github.com/mono83/cfg/file"
	"github.com/mono83/cfg/json"
)

// JSONFile registers configuration source read from JSON file
func JSONFile(names ...string) {
	def.JSONFile(names...)
}

// JSONFile registers configuration source read from JSON file
func (c *Config) JSONFile(names ...string) {
	c.AddLast(file.New(names, nil, json.NewBytesSource))
}

// JSONAutoFind registers configuration source read from required file
// File will searched in current folder, then in home folder and then in /etc/
func JSONAutoFind(filename, subfolderName string) {
	def.JSONAutoFind(filename, subfolderName)
}

// JSONAutoFind registers configuration source read from required file
// File will searched in current folder, then in home folder and then in /etc/
func (c *Config) JSONAutoFind(filename, subfolderName string) {
	if subfolderName == "" {
		c.JSONFile(file.CommonFolders(filename)...)
	} else {
		c.JSONFile(file.CommonFoldersWithSubfolder(filename, subfolderName)...)
	}
}
