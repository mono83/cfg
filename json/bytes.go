package json

import (
	"encoding/json"
	"github.com/mono83/cfg"
	"io"
	"io/ioutil"
)

// NewReaderSource reads all bytes from source reader and runs NewBytesSource
// creating configurer from JSON bytes
func NewReaderSource(source io.Reader) (cfg.Configurer, error) {
	bts, err := ioutil.ReadAll(source)
	if err != nil {
		return nil, err
	}

	return NewBytesSource(bts)
}

// NewStringSource converts source to []byte and runs NewBytesSource
// creating configurer from JSON bytes
func NewStringSource(source string) (cfg.Configurer, error) {
	return NewBytesSource([]byte(source))
}

// NewBytesSource creates and returns JSON configurer, that uses
// provided bytes as source data
func NewBytesSource(source []byte) (cfg.Configurer, error) {
	var t map[string]json.RawMessage
	err := json.Unmarshal(source, &t)
	if err != nil {
		return nil, err
	}

	return bytesSourceConfigurer(t), nil
}

type bytesSourceConfigurer map[string]json.RawMessage

func (b bytesSourceConfigurer) Has(key string) bool {
	_, ok := b[key]
	return ok
}

func (b bytesSourceConfigurer) UnmarshallKey(key string, target interface{}) error {
	v, ok := b[key]
	if !ok {
		return cfg.ErrKeyMissing{Key: key}
	}

	return json.Unmarshal(v, target)
}

func (b bytesSourceConfigurer) KeyFunc(key string) func(interface{}) error {
	return cfg.ExtractUnmarshallFunc(b, key)
}
