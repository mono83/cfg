package yaml

import (
	"fmt"
	"github.com/mono83/cfg"
	"gopkg.in/yaml.v2"
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
	var target yaml.MapSlice
	err := yaml.Unmarshal(source, &target)
	if err != nil {
		return nil, err
	}

	// Converting to map[string]interface{}
	response := map[string][]byte{}
	for _, item := range target {
		out, err := yaml.Marshal(item.Value)
		if err != nil {
			return nil, err
		}
		if lo := len(out); lo > 0 && out[lo-1] == 10 {
			out = out[0 : lo-1]
		}
		response[fmt.Sprintf("%v", item.Key)] = out
	}

	return bytesSourceConfigurer(response), nil
}

type bytesSourceConfigurer map[string][]byte

func (b bytesSourceConfigurer) Has(key string) bool {
	_, ok := b[key]
	return ok
}

func (b bytesSourceConfigurer) UnmarshallKey(key string, target interface{}) error {
	v, ok := b[key]
	if !ok {
		return cfg.ErrKeyMissing{Key: key}
	}

	return yaml.Unmarshal(v, target)
}

func (b bytesSourceConfigurer) KeyFunc(key string) func(interface{}) error {
	return cfg.ExtractUnmarshallFunc(b, key)
}
