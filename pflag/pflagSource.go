package pflag

import (
	"github.com/mono83/cfg"
	"github.com/mono83/cfg/reflect"
	"github.com/ogier/pflag"
	"os"
	"sync"
)

type flagSource struct {
	set  *pflag.FlagSet
	args []string
	m    sync.Mutex

	values map[string]interface{}
}

// NewPFlagSource creates new configuration source from command line flags
func NewPFlagSource() cfg.Configurer {
	return NewCustomPFlagSource(pflag.CommandLine, os.Args[1:])
}

// NewCustomPFlagSource creates new configuration source from provided
// FlagSet and args
func NewCustomPFlagSource(source *pflag.FlagSet, args []string) cfg.Configurer {
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

	f.set.Visit(func(fl *pflag.Flag) {
		f.values[fl.Name] = fl.Value.String()
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
