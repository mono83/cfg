package file

import (
	"github.com/mono83/cfg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlaceholdersReaderSuccess(t *testing.T) {
	a := assert.New(t)

	source := func(string) ([]byte, error) {
		return []byte("something: \"#var#\"\nmore: #var2#"), nil
	}

	values := cfg.Map(map[string]interface{}{
		"var":  "foo",
		"var2": 0.21,
	})

	phrf := PlaceholdersReader(source, values)

	data, err := phrf("ignore")
	a.NoError(err)
	a.NotNil(data)
	a.Equal("something: \"foo\"\nmore: 0.21", string(data))
}
