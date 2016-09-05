package cfg

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	a := assert.New(t)

	first := Map(map[string]interface{}{
		"foo": 10,
		"bar": 20,
		"baz": 30,
	})
	second := Map(map[string]interface{}{
		"foo2": 110,
		"bar":  120,
		"baz2": 130,
	})

	l := List([]Configurer{first, second})

	a.NoError(l.Validate())

	a.True(l.Has("foo"))
	a.True(l.Has("bar"))
	a.True(l.Has("baz"))
	a.True(l.Has("foo2"))
	a.True(l.Has("baz2"))

	i := 0
	a.NoError(l.UnmarshalKey("foo", &i))
	a.Equal(10, i)
	a.NoError(l.UnmarshalKey("bar", &i))
	a.Equal(120, i)
	a.NoError(l.UnmarshalKey("baz", &i))
	a.Equal(30, i)
	a.NoError(l.UnmarshalKey("foo2", &i))
	a.Equal(110, i)
	a.NoError(l.UnmarshalKey("baz2", &i))
	a.Equal(130, i)

	err := errors.New("Marker error")
	l = List([]Configurer{first, ErrConfigurer(err), second})
	a.Error(l.Validate())
	a.Equal(err.Error(), l.Validate().Error())

	// Even validation failure of one of configurers allows to use others
	a.True(l.Has("foo"))
	a.True(l.Has("bar"))
	a.True(l.Has("baz"))
	a.True(l.Has("foo2"))
	a.True(l.Has("baz2"))
}
