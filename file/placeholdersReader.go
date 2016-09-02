package file

import (
	"fmt"
	"github.com/mono83/cfg"
	"regexp"
)

var placeholdersRegex = regexp.MustCompile("(#[a-zA-Z\\.\\-0-9]+#)")

// PlaceholdersReader return file reader function, that reads file
// from filesystem and replaces all placeholders in it using
// previously loaded configs
// Ordering matters
func PlaceholdersReader(read func(string) ([]byte, error), configurer cfg.Configurer) func(string) ([]byte, error) {
	return func(name string) ([]byte, error) {
		// Reading using external reader
		data, err := read(name)
		if err != nil {
			return nil, err
		}

		// Got some data - searching and replacing placeholders
		var outerError error
		bts := placeholdersRegex.ReplaceAllFunc(data, func(b []byte) []byte {
			if outerError != nil {
				return b
			}

			s := string(b[1 : len(b)-1])

			// Searching for key
			if !configurer.Has(s) {
				outerError = fmt.Errorf("Configuration has no data for placeholder %s", s)
				return b
			}

			var target interface{}
			outerError = configurer.UnmarshalKey(s, &target)
			if outerError == nil {
				b = []byte(fmt.Sprintf("%v", target))
			}

			return b
		})

		return bts, outerError
	}
}
