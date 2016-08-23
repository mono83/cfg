package cfg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigurerWithAliases(t *testing.T) {
	a := assert.New(t)

	ms := map[string]interface{}{
		"foo": 0.231,
	}

	f := .0

	c := NewConfigurationWithAliases(Map(ms))
	a.True(c.Has("foo"))
	a.False(c.Has("bar"))
	a.NoError(c.UnmarshallKey("foo", &f))
	a.Error(c.UnmarshallKey("bar", &f))

	c.Alias("bar", "foo")
	a.True(c.Has("foo"))
	a.True(c.Has("bar"))
	a.NoError(c.UnmarshallKey("foo", &f))
	a.NoError(c.UnmarshallKey("bar", &f))
	a.Equal(0.231, f)
}
