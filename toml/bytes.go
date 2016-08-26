package toml

import (
	"github.com/BurntSushi/toml"
	"github.com/mono83/cfg"
	"io"
	"io/ioutil"
)

// NewReaderSource reads all bytes from source reader and runs NewBytesSource
// creating configurer from YAML bytes
func NewReaderSource(source io.Reader) (cfg.Configurer, error) {
	bts, err := ioutil.ReadAll(source)
	if err != nil {
		return nil, err
	}

	return NewBytesSource(bts)
}

// NewStringSource converts source to []byte and runs NewBytesSource
// creating configurer from YAML bytes
func NewStringSource(source string) (cfg.Configurer, error) {
	return NewBytesSource([]byte(source))
}

// NewBytesSource creates and returns YAML configurer, that uses
// provided bytes as source data
func NewBytesSource(source []byte) (cfg.Configurer, error) {
	var target map[string]toml.Primitive
	err := toml.Unmarshal(source, &target)
	if err != nil {
		return nil, err
	}

	return bytesSourceConfigurer(target), nil
}

type bytesSourceConfigurer map[string]toml.Primitive

func (b bytesSourceConfigurer) Has(key string) bool {
	_, ok := b[key]
	return ok
}

func (b bytesSourceConfigurer) UnmarshalKey(key string, target interface{}) error {
	v, ok := b[key]
	if !ok {
		return cfg.ErrKeyMissing{Key: key}
	}

	return toml.PrimitiveDecode(v, target)
}

func (b bytesSourceConfigurer) KeyFunc(key string) func(interface{}) error {
	return cfg.ExtractUnmarshalFunc(b, key)
}
