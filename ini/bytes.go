package ini

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/mono83/cfg"
	"io"
	"io/ioutil"
	"strings"
)

var nlByte = []byte{0x0a}

// NewReaderSource reads all bytes from source reader and runs NewBytesSource
// creating configurer from INI file bytes
func NewReaderSource(source io.Reader) (cfg.Configurer, error) {
	bts, err := ioutil.ReadAll(source)
	if err != nil {
		return nil, err
	}

	return NewBytesSource(bts)
}

// NewStringSource converts source to []byte and runs NewBytesSource
// creating configurer from INI file bytes
func NewStringSource(source string) (cfg.Configurer, error) {
	return NewBytesSource([]byte(source))
}

// NewBytesSource creates and returns map configurer, that uses
// provided bytes as source data
func NewBytesSource(source []byte) (cfg.Configurer, error) {
	if len(source) == 0 {
		return nil, errors.New("Empty INI source")
	}

	src := map[string]interface{}{}
	for _, bline := range bytes.Split(source, nlByte) {
		line := strings.TrimSpace(string(bline))
		if len(line) == 0 || line[0] == '#' || line[0] == '/' || line[0] == '[' {
			// Skipping
			continue
		}

		pos := strings.IndexRune(line, '=')
		if pos == -1 {
			return nil, fmt.Errorf("Unable to parse line %s", line)
		}

		key := strings.TrimSpace(line[0:pos])
		value := strings.TrimSpace(line[pos+1:])

		src[key] = value
	}

	return cfg.Map(src), nil
}
