package config

import (
	"github.com/mono83/cfg/file"
	"github.com/mono83/cfg/ini"
)

// INIFile registers configuration source read from JSON file
func INIFile(names ...string) {
	def.INIFile(names...)
}

// INIFile registers configuration source read from JSON file
func (c *Config) INIFile(names ...string) {
	c.AddLast(file.New(names, c.fileReader, ini.NewBytesSource))
}

// INIAutoFind registers configuration source read from required file
// File will searched in current folder, then in home folder and then in /etc/
func INIAutoFind(filename, subfolderName string) {
	def.INIAutoFind(filename, subfolderName)
}

// INIAutoFind registers configuration source read from required file
// File will searched in current folder, then in home folder and then in /etc/
func (c *Config) INIAutoFind(filename, subfolderName string) {
	if subfolderName == "" {
		c.INIFile(file.CommonFolders(filename)...)
	} else {
		c.INIFile(file.CommonFoldersWithSubfolder(filename, subfolderName)...)
	}
}
