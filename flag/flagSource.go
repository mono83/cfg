package flag

import (
	"flag"
	"github.com/mono83/cfg"
	"github.com/mono83/cfg/reflect"
	"os"
	"sync"
)

type flagSource struct {
	set  *flag.FlagSet
	args []string
	m    sync.Mutex

	values map[string]interface{}
}

// NewFlagSource creates new configuration source from command line flags
func NewFlagSource() cfg.Configurer {
	return NewCustomFlagSource(flag.CommandLine, os.Args[1:])
}

// NewCustomFlagSource creates new configuration source from provided
// FlagSet and args
func NewCustomFlagSource(source *flag.FlagSet, args []string) cfg.Configurer {
	return &flagSource{set: source, args: args}
}

func (f *flagSource) Validate() error {
	return f.load()
}

func (f *flagSource) load() error {
	f.m.Lock()
	defer f.m.Unlock()

	if !f.set.Parsed() {
		err := f.set.Parse(f.args)
		if err != nil {
			return err
		}
	}

	f.values = map[string]interface{}{}

	f.set.Visit(func(fl *flag.Flag) {
		v, ok := fl.Value.(flag.Getter)
		if ok {
			f.values[fl.Name] = v.Get()
		}
	})

	return nil
}

func (f *flagSource) Has(key string) bool {
	if f.values == nil {
		err := f.load()
		if err != nil {
			return false
		}
	}

	_, ok := f.values[key]
	return ok
}

func (f *flagSource) UnmarshalKey(key string, target interface{}) error {
	if f.values == nil {
		err := f.load()
		if err != nil {
			return err
		}
	}

	v, ok := f.values[key]
	if !ok {
		return cfg.ErrKeyMissing{Key: key}
	}

	return reflect.CopyHelper(key, v, target)
}

func (f *flagSource) KeyFunc(key string) func(interface{}) error {
	return cfg.ExtractUnmarshalFunc(f, key)
}
