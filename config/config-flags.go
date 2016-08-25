package config

import (
	goflag "flag"
	"github.com/mono83/cfg/flag"
	"github.com/mono83/cfg/pflag"
	gopflag "github.com/ogier/pflag"
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

// EnablePFlags enables improved flags reading
func EnablePFlags() {
	def.EnablePFlags()
}

// EnablePFlags enables improved flags reading
func (c *Config) EnablePFlags() {
	c.AddLast(pflag.NewPFlagSource())
}

// EnableCustomPFlags registers configuration source from
// provided improved PFlagSet
func EnableCustomPFlags(f *gopflag.FlagSet) {
	def.EnableCustomFlags(f)
}

// EnableCustomPFlags registers configuration source from
// provided improved PFlagSet
func (c *Config) EnableCustomPFlags(f *gopflag.FlagSet) {
	c.AddLast(pflag.NewCustomPFlagSource(f, os.Args[1:]))
}
