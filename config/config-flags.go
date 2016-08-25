package config

import (
	goflag "flag"
	"github.com/mono83/cfg/flag"
	"os"
)

// EnableFlags enables standard library flags reading
func EnableFlags() {
	def.EnableFlags()
}

// EnableFlags enables standard library flags reading
func (c *Config) EnableFlags() {
	c.AddLast(flag.NewFlagSource())
}

// EnableCustomFlags registers configuration source from
// provided FlagSet
func EnableCustomFlags(f *goflag.FlagSet) {
	def.EnableCustomFlags(f)
}

// EnableCustomFlags registers configuration source from
// provided FlagSet
func (c *Config) EnableCustomFlags(f *goflag.FlagSet) {
	c.AddLast(flag.NewCustomFlagSource(f, os.Args[1:]))
}
