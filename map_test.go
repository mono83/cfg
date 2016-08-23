package cfg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testMapEntry struct {
	id   int
	name string
}

func TestMap(t *testing.T) {
	a := assert.New(t)

	ms := map[string]interface{}{
		"foo":   "bar",
		"id":    462,
		"entry": testMapEntry{id: 10, name: "hello, world"},
	}

	m := Map(ms)
	a.True(m.Has("foo"))
	a.True(m.Has("id"))
	a.False(m.Has("Foo"))

	s := ""
	a.NoError(m.UnmarshalKey("foo", &s))
	a.Equal("bar", s)
	i := 0
	a.NoError(m.UnmarshalKey("id", &i))
	a.Equal(462, i)

	tme := testMapEntry{}
	a.NoError(m.UnmarshalKey("entry", &tme))
	a.Equal(10, tme.id)
	a.Equal("hello, world", tme.name)
}
