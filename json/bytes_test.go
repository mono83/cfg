package json

import (
	"github.com/mono83/cfg"
	"github.com/stretchr/testify/assert"
	"testing"
)

var jsonSrc = `
{
   "name": "foo",
   "id": 124,
   "amount": 32.55,
   "enabled": true
}`

func TestInvalidJSON(t *testing.T) {
	invalid := `{"name":"foo"`
	r, err := NewBytesSource([]byte(invalid))

	assert.Nil(t, r)
	assert.Error(t, err)
}

func TestUnmarshall(t *testing.T) {
	a := assert.New(t)

	c, err := NewBytesSource([]byte(jsonSrc))

	a.NoError(err)
	a.True(c.Has("name"))
	a.True(c.Has("id"))
	a.True(c.Has("amount"))
	a.True(c.Has("enabled"))
	a.False(c.Has("Name"))
	a.False(c.Has("ID"))
	a.False(c.Has("Amount"))
	a.False(c.Has("Enabled"))

	var s string
	a.NoError(c.UnmarshallKey("name", &s))
	a.Equal("foo", s)
	s = ""
	a.NoError(c.KeyFunc("name")(&s))
	a.Equal("foo", s)

	var i int
	a.NoError(c.UnmarshallKey("id", &i))
	a.Equal(124, i)

	var f float64
	a.NoError(c.UnmarshallKey("amount", &f))
	a.Equal(32.55, f)

	var b bool
	a.NoError(c.UnmarshallKey("enabled", &b))
	a.Equal(true, b)

	err = c.UnmarshallKey("missing", &b)
	a.True(cfg.IsErrKeyMissing(err))
}
