package yaml

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var yamlSrc = `
stringKey: "some"
stringKey2: more
intKey: 736
boolKey: true
floatKey: 0.323
`

func TestBytesSourceConfigurer(t *testing.T) {
	a := assert.New(t)

	c, err := NewBytesSource([]byte(yamlSrc))
	a.NoError(err)

	a.True(c.Has("stringKey"))
	a.True(c.Has("stringKey2"))
	a.True(c.Has("intKey"))
	a.True(c.Has("boolKey"))
	a.True(c.Has("floatKey"))
	a.False(c.Has("floatKEY"))

	s := ""
	a.NoError(c.UnmarshallKey("stringKey", &s))
	a.Equal("some", s)
	a.NoError(c.UnmarshallKey("stringKey2", &s))
	a.Equal("more", s)

	i := 0
	a.NoError(c.UnmarshallKey("intKey", &i))
	a.Equal(736, i)

	b := false
	a.NoError(c.UnmarshallKey("boolKey", &b))
	a.True(b)

	f := 0.0
	a.NoError(c.UnmarshallKey("floatKey", &f))
	a.Equal(0.323, f)
}
